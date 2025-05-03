package particles

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// HugoShortcodeData is passed to the Hugo shortcode
type HugoShortcodeData struct {
	ElementID      string
	ConfigEndpoint string
	JsPath         string
}

// HugoHandler processes particles requests for Hugo
type HugoHandler struct {
	ConfigEndpoint  string
	StaticJsPath    string
	ParticlesCache  map[string]*Config
	DefaultConfigID string
}

// NewHugoHandler creates a new Hugo handler
func NewHugoHandler(configEndpoint, staticJsPath string) *HugoHandler {
	// Generate a default config ID
	defaultID := fmt.Sprintf("config-%d", time.Now().UnixNano())

	handler := &HugoHandler{
		ConfigEndpoint:  configEndpoint,
		StaticJsPath:    staticJsPath,
		ParticlesCache:  make(map[string]*Config),
		DefaultConfigID: defaultID,
	}

	// Add default config to cache
	handler.ParticlesCache[defaultID] = DefaultConfig()

	return handler
}

// ServeHTTP handles HTTP requests for particle configurations
func (h *HugoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Set JSON content type
	w.Header().Set("Content-Type", "application/json")

	// Parse request
	configID := r.URL.Query().Get("config")
	if configID == "" {
		configID = h.DefaultConfigID
	}

	// Check if we already have this config
	config, exists := h.ParticlesCache[configID]
	if !exists {
		// Generate a random config
		config = RandomParticlesConfig()
		h.ParticlesCache[configID] = config
	}

	// Marshal config to JSON
	jsonData, err := json.Marshal(config)
	if err != nil {
		http.Error(w, "Error generating JSON", http.StatusInternalServerError)
		return
	}

	// Write JSON response
	w.Write(jsonData)
}

// GenerateHugoShortcodeData creates data for the Hugo shortcode
func (h *HugoHandler) GenerateHugoShortcodeData(params map[string]string) HugoShortcodeData {
	// Get or create a config ID
	configID := params["config"]
	if configID == "" {
		configID = fmt.Sprintf("particles-%d", rand.Int63())
	}

	// Create a specific element ID
	elementID := params["id"]
	if elementID == "" {
		elementID = fmt.Sprintf("particles-%s", configID)
	}

	// Create a configuration for this instance
	config := DefaultConfig()

	// Process parameters to update config
	for k, v := range params {
		switch k {
		case "preset":
			if preset, ok := v, true; ok {
				config = GetPreset(preset)
			}
		case "color":
			config.Particles.Color.Value = v
		case "shape":
			config.Particles.Shape.Type = v
		case "number":
			if val, err := strconv.Atoi(v); err == nil {
				config.Particles.Number.Value = val
			}
		case "size":
			if val, err := strconv.ParseFloat(v, 64); err == nil {
				config.Particles.Size.Value = val
			}
		case "speed":
			if val, err := strconv.ParseFloat(v, 64); err == nil {
				config.Particles.Move.Speed = val
			}
		case "direction":
			config.Particles.Move.Direction = v
		}
	}

	// Store config in cache
	h.ParticlesCache[configID] = config

	// Build config endpoint URL
	configURL := fmt.Sprintf("%s?config=%s", h.ConfigEndpoint, configID)

	return HugoShortcodeData{
		ElementID:      elementID,
		ConfigEndpoint: configURL,
		JsPath:         h.StaticJsPath,
	}
}

// Shortcode generates the HTML for the Hugo shortcode
func (h *HugoHandler) Shortcode(params map[string]string) template.HTML {
	data := h.GenerateHugoShortcodeData(params)

	// Create HTML output
	html := fmt.Sprintf(`
<div id="%s" style="width: 100%%; height: 100%%; position: absolute; top: 0; left: 0; z-index: -1;"></div>
<script src="%s"></script>
<script>
document.addEventListener('DOMContentLoaded', function() {
  particlesJS.load('%s', '%s', function() {
    console.log('particles.js loaded');
  });
});
</script>
`, data.ElementID, data.JsPath, data.ElementID, data.ConfigEndpoint)

	return template.HTML(html)
}
