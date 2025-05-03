package particles

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Color represents a color in RGB or HSL format
type Color struct {
	Value interface{} `json:"value"`
}

// NumberDensity represents the number density configuration
type NumberDensity struct {
	Enable    bool    `json:"enable"`
	ValueArea float64 `json:"value_area"`
}

// Number represents the number configuration
type Number struct {
	Value   int           `json:"value"`
	Density NumberDensity `json:"density"`
}

// ShapeStroke represents the shape stroke configuration
type ShapeStroke struct {
	Width float64 `json:"width"`
	Color string  `json:"color"`
}

// ShapePolygon represents the polygon shape configuration
type ShapePolygon struct {
	NbSides int `json:"nb_sides"`
}

// ShapeImage represents the image shape configuration
type ShapeImage struct {
	Src    string  `json:"src"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

// Shape represents the shape configuration
type Shape struct {
	Type    interface{}  `json:"type"`
	Stroke  ShapeStroke  `json:"stroke"`
	Polygon ShapePolygon `json:"polygon"`
	Image   ShapeImage   `json:"image"`
}

// OpacityAnimation represents the opacity animation configuration
type OpacityAnimation struct {
	Enable     bool    `json:"enable"`
	Speed      float64 `json:"speed"`
	OpacityMin float64 `json:"opacity_min"`
	Sync       bool    `json:"sync"`
}

// Opacity represents the opacity configuration
type Opacity struct {
	Value  float64          `json:"value"`
	Random bool             `json:"random"`
	Anim   OpacityAnimation `json:"anim"`
}

// SizeAnimation represents the size animation configuration
type SizeAnimation struct {
	Enable  bool    `json:"enable"`
	Speed   float64 `json:"speed"`
	SizeMin float64 `json:"size_min"`
	Sync    bool    `json:"sync"`
}

// Size represents the size configuration
type Size struct {
	Value  float64       `json:"value"`
	Random bool          `json:"random"`
	Anim   SizeAnimation `json:"anim"`
}

// LineLinked represents the line linked configuration
type LineLinked struct {
	Enable   bool    `json:"enable"`
	Distance float64 `json:"distance"`
	Color    string  `json:"color"`
	Opacity  float64 `json:"opacity"`
	Width    float64 `json:"width"`
}

// MoveAttract represents the move attract configuration
type MoveAttract struct {
	Enable  bool    `json:"enable"`
	RotateX float64 `json:"rotateX"`
	RotateY float64 `json:"rotateY"`
}

// Move represents the move configuration
type Move struct {
	Enable    bool        `json:"enable"`
	Speed     float64     `json:"speed"`
	Direction string      `json:"direction"`
	Random    bool        `json:"random"`
	Straight  bool        `json:"straight"`
	OutMode   string      `json:"out_mode"`
	Bounce    bool        `json:"bounce"`
	Attract   MoveAttract `json:"attract"`
}

// ParticlesConfig represents the particles configuration
type ParticlesConfig struct {
	Number     Number     `json:"number"`
	Color      Color      `json:"color"`
	Shape      Shape      `json:"shape"`
	Opacity    Opacity    `json:"opacity"`
	Size       Size       `json:"size"`
	LineLinked LineLinked `json:"line_linked"`
	Move       Move       `json:"move"`
}

// InteractivityEventMode represents the interactivity event mode
type InteractivityEventMode struct {
	Enable bool   `json:"enable"`
	Mode   string `json:"mode"`
}

// InteractivityEvents represents the interactivity events
type InteractivityEvents struct {
	OnHover InteractivityEventMode `json:"onhover"`
	OnClick InteractivityEventMode `json:"onclick"`
	Resize  bool                   `json:"resize"`
}

// GrabLineLinked represents the grab line linked configuration
type GrabLineLinked struct {
	Opacity float64 `json:"opacity"`
}

// GrabMode represents the grab mode configuration
type GrabMode struct {
	Distance   float64        `json:"distance"`
	LineLinked GrabLineLinked `json:"line_linked"`
}

// BubbleMode represents the bubble mode configuration
type BubbleMode struct {
	Distance float64 `json:"distance"`
	Size     float64 `json:"size"`
	Duration float64 `json:"duration"`
	Opacity  float64 `json:"opacity"`
	Speed    float64 `json:"speed"`
}

// RepulseMode represents the repulse mode configuration
type RepulseMode struct {
	Distance float64 `json:"distance"`
	Duration float64 `json:"duration"`
}

// PushMode represents the push mode configuration
type PushMode struct {
	ParticlesNb int `json:"particles_nb"`
}

// RemoveMode represents the remove mode configuration
type RemoveMode struct {
	ParticlesNb int `json:"particles_nb"`
}

// InteractivityModes represents the interactivity modes
type InteractivityModes struct {
	Grab    GrabMode    `json:"grab"`
	Bubble  BubbleMode  `json:"bubble"`
	Repulse RepulseMode `json:"repulse"`
	Push    PushMode    `json:"push"`
	Remove  RemoveMode  `json:"remove"`
}

// Interactivity represents the interactivity configuration
type Interactivity struct {
	DetectOn string              `json:"detect_on"`
	Events   InteractivityEvents `json:"events"`
	Modes    InteractivityModes  `json:"modes"`
}

// Config represents the overall particles.js configuration
type Config struct {
	Particles     ParticlesConfig `json:"particles"`
	Interactivity Interactivity   `json:"interactivity"`
	RetinaDetect  bool            `json:"retina_detect"`
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		Particles: ParticlesConfig{
			Number: Number{
				Value: 80,
				Density: NumberDensity{
					Enable:    true,
					ValueArea: 800,
				},
			},
			Color: Color{
				Value: "#ffffff",
			},
			Shape: Shape{
				Type: "circle",
				Stroke: ShapeStroke{
					Width: 0,
					Color: "#000000",
				},
				Polygon: ShapePolygon{
					NbSides: 5,
				},
				Image: ShapeImage{
					Src:    "",
					Width:  100,
					Height: 100,
				},
			},
			Opacity: Opacity{
				Value:  0.5,
				Random: false,
				Anim: OpacityAnimation{
					Enable:     false,
					Speed:      1,
					OpacityMin: 0.1,
					Sync:       false,
				},
			},
			Size: Size{
				Value:  5,
				Random: true,
				Anim: SizeAnimation{
					Enable:  false,
					Speed:   40,
					SizeMin: 0.1,
					Sync:    false,
				},
			},
			LineLinked: LineLinked{
				Enable:   true,
				Distance: 150,
				Color:    "#ffffff",
				Opacity:  0.4,
				Width:    1,
			},
			Move: Move{
				Enable:    true,
				Speed:     6,
				Direction: "none",
				Random:    false,
				Straight:  false,
				OutMode:   "out",
				Bounce:    false,
				Attract: MoveAttract{
					Enable:  false,
					RotateX: 600,
					RotateY: 1200,
				},
			},
		},
		Interactivity: Interactivity{
			DetectOn: "canvas",
			Events: InteractivityEvents{
				OnHover: InteractivityEventMode{
					Enable: true,
					Mode:   "repulse",
				},
				OnClick: InteractivityEventMode{
					Enable: true,
					Mode:   "push",
				},
				Resize: true,
			},
			Modes: InteractivityModes{
				Grab: GrabMode{
					Distance: 400,
					LineLinked: GrabLineLinked{
						Opacity: 1,
					},
				},
				Bubble: BubbleMode{
					Distance: 400,
					Size:     40,
					Duration: 2,
					Opacity:  8,
					Speed:    3,
				},
				Repulse: RepulseMode{
					Distance: 200,
					Duration: 0.4,
				},
				Push: PushMode{
					ParticlesNb: 4,
				},
				Remove: RemoveMode{
					ParticlesNb: 2,
				},
			},
		},
		RetinaDetect: true,
	}
}

// Preset defined particle configurations
const (
	PresetDefault   = "default"
	PresetSnow      = "snow"
	PresetNightSky  = "nightsky"
	PresetSpacyDots = "spacydots"
	PresetBubbles   = "bubbles"
)

// GetPreset returns a predefined configuration
func GetPreset(preset string) *Config {
	switch preset {
	case PresetSnow:
		return &Config{
			Particles: ParticlesConfig{
				Number: Number{
					Value: 400,
					Density: NumberDensity{
						Enable:    true,
						ValueArea: 800,
					},
				},
				Color: Color{
					Value: "#ffffff",
				},
				Shape: Shape{
					Type: "circle",
					Stroke: ShapeStroke{
						Width: 0,
						Color: "#000000",
					},
					Polygon: ShapePolygon{
						NbSides: 5,
					},
				},
				Opacity: Opacity{
					Value:  0.5,
					Random: true,
				},
				Size: Size{
					Value:  3,
					Random: true,
				},
				LineLinked: LineLinked{
					Enable: false,
				},
				Move: Move{
					Enable:    true,
					Speed:     2,
					Direction: "bottom",
					Random:    true,
					Straight:  false,
					OutMode:   "out",
					Bounce:    false,
				},
			},
			Interactivity: Interactivity{
				DetectOn: "canvas",
				Events: InteractivityEvents{
					OnHover: InteractivityEventMode{
						Enable: false,
					},
					OnClick: InteractivityEventMode{
						Enable: true,
						Mode:   "repulse",
					},
					Resize: true,
				},
			},
			RetinaDetect: true,
		}
	case PresetNightSky:
		return &Config{
			Particles: ParticlesConfig{
				Number: Number{
					Value: 160,
					Density: NumberDensity{
						Enable:    true,
						ValueArea: 800,
					},
				},
				Color: Color{
					Value: "#ffffff",
				},
				Shape: Shape{
					Type: "circle",
				},
				Opacity: Opacity{
					Value:  0.8,
					Random: true,
					Anim: OpacityAnimation{
						Enable:     true,
						Speed:      1,
						OpacityMin: 0.1,
						Sync:       false,
					},
				},
				Size: Size{
					Value:  3,
					Random: true,
				},
				LineLinked: LineLinked{
					Enable:   true,
					Distance: 100,
					Color:    "#ffffff",
					Opacity:  0.2,
					Width:    1,
				},
				Move: Move{
					Enable:    true,
					Speed:     1,
					Direction: "none",
					Random:    true,
				},
			},
			Interactivity: Interactivity{
				DetectOn: "canvas",
				Events: InteractivityEvents{
					OnHover: InteractivityEventMode{
						Enable: true,
						Mode:   "bubble",
					},
					OnClick: InteractivityEventMode{
						Enable: true,
						Mode:   "push",
					},
				},
				Modes: InteractivityModes{
					Bubble: BubbleMode{
						Distance: 250,
						Size:     5,
						Duration: 2,
					},
				},
			},
		}
	case PresetSpacyDots:
		return &Config{
			Particles: ParticlesConfig{
				Number: Number{
					Value: 120,
					Density: NumberDensity{
						Enable:    true,
						ValueArea: 800,
					},
				},
				Color: Color{
					Value: "#ffffff",
				},
				Shape: Shape{
					Type: "circle",
				},
				Opacity: Opacity{
					Value:  0.5,
					Random: false,
				},
				Size: Size{
					Value:  3,
					Random: true,
				},
				LineLinked: LineLinked{
					Enable:   true,
					Distance: 150,
					Color:    "#ffffff",
					Opacity:  0.4,
					Width:    1,
				},
				Move: Move{
					Enable:    true,
					Speed:     3,
					Direction: "none",
					Random:    false,
					Straight:  false,
					OutMode:   "out",
					Bounce:    false,
					Attract: MoveAttract{
						Enable:  false,
						RotateX: 600,
						RotateY: 1200,
					},
				},
			},
			Interactivity: Interactivity{
				DetectOn: "canvas",
				Events: InteractivityEvents{
					OnHover: InteractivityEventMode{
						Enable: true,
						Mode:   "grab",
					},
					OnClick: InteractivityEventMode{
						Enable: true,
						Mode:   "push",
					},
					Resize: true,
				},
				Modes: InteractivityModes{
					Grab: GrabMode{
						Distance: 140,
						LineLinked: GrabLineLinked{
							Opacity: 1,
						},
					},
					Push: PushMode{
						ParticlesNb: 4,
					},
				},
			},
			RetinaDetect: true,
		}
	case PresetBubbles:
		return &Config{
			Particles: ParticlesConfig{
				Number: Number{
					Value: 50,
					Density: NumberDensity{
						Enable:    true,
						ValueArea: 800,
					},
				},
				Color: Color{
					Value: "#4285f4",
				},
				Shape: Shape{
					Type: "circle",
					Stroke: ShapeStroke{
						Width: 0,
						Color: "#000000",
					},
				},
				Opacity: Opacity{
					Value:  0.5,
					Random: true,
					Anim: OpacityAnimation{
						Enable:     true,
						Speed:      3,
						OpacityMin: 0.1,
						Sync:       false,
					},
				},
				Size: Size{
					Value:  15,
					Random: true,
					Anim: SizeAnimation{
						Enable:  true,
						Speed:   5,
						SizeMin: 0.1,
						Sync:    false,
					},
				},
				LineLinked: LineLinked{
					Enable: false,
				},
				Move: Move{
					Enable:    true,
					Speed:     3,
					Direction: "none",
					Random:    true,
					Straight:  false,
					OutMode:   "out",
					Bounce:    false,
				},
			},
			Interactivity: Interactivity{
				DetectOn: "canvas",
				Events: InteractivityEvents{
					OnHover: InteractivityEventMode{
						Enable: true,
						Mode:   "bubble",
					},
					OnClick: InteractivityEventMode{
						Enable: true,
						Mode:   "repulse",
					},
				},
				Modes: InteractivityModes{
					Bubble: BubbleMode{
						Distance: 250,
						Size:     0,
						Duration: 2,
						Opacity:  0,
						Speed:    3,
					},
					Repulse: RepulseMode{
						Distance: 400,
						Duration: 0.4,
					},
				},
			},
			RetinaDetect: true,
		}
	default:
		return DefaultConfig()
	}
}

// GenerateConfig creates a particle configuration with the given parameters
func GenerateConfig(params map[string]interface{}) *Config {
	// Start with default config
	config := DefaultConfig()

	// Check if a preset was specified
	if presetStr, ok := params["preset"].(string); ok && presetStr != "" {
		config = GetPreset(presetStr)
	}

	// Override with any passed parameters
	if number, ok := params["number"].(int); ok {
		config.Particles.Number.Value = number
	}

	if color, ok := params["color"].(string); ok {
		config.Particles.Color.Value = color
	}

	if shape, ok := params["shape"].(string); ok {
		config.Particles.Shape.Type = shape
	}

	if size, ok := params["size"].(float64); ok {
		config.Particles.Size.Value = size
	}

	if speed, ok := params["speed"].(float64); ok {
		config.Particles.Move.Speed = speed
	}

	if direction, ok := params["direction"].(string); ok {
		config.Particles.Move.Direction = direction
	}

	if opacity, ok := params["opacity"].(float64); ok {
		config.Particles.Opacity.Value = opacity
	}

	if lineColor, ok := params["lineColor"].(string); ok {
		config.Particles.LineLinked.Color = lineColor
	}

	if lineWidth, ok := params["lineWidth"].(float64); ok {
		config.Particles.LineLinked.Width = lineWidth
	}

	if lineDistance, ok := params["lineDistance"].(float64); ok {
		config.Particles.LineLinked.Distance = lineDistance
	}

	if hoverMode, ok := params["hoverMode"].(string); ok {
		config.Interactivity.Events.OnHover.Mode = hoverMode
	}

	if clickMode, ok := params["clickMode"].(string); ok {
		config.Interactivity.Events.OnClick.Mode = clickMode
	}

	return config
}

// ToJSON converts the configuration to a JSON string
func (c *Config) ToJSON() (string, error) {
	bytes, err := json.Marshal(c)
	if err != nil {
		return "", fmt.Errorf("error marshaling config to JSON: %v", err)
	}
	return string(bytes), nil
}

// RandomParticlesConfig generates a random particles configuration
func RandomParticlesConfig() *Config {
	config := DefaultConfig()

	// Randomize particles number
	config.Particles.Number.Value = rand.Intn(150) + 50

	// Randomize particles color
	colors := []string{"#ffffff", "#e74c3c", "#3498db", "#2ecc71", "#f1c40f", "#9b59b6"}
	config.Particles.Color.Value = colors[rand.Intn(len(colors))]

	// Randomize particles shape
	shapes := []string{"circle", "edge", "triangle", "polygon", "star"}
	config.Particles.Shape.Type = shapes[rand.Intn(len(shapes))]

	// Randomize particles size
	config.Particles.Size.Value = float64(rand.Intn(10) + 1)
	config.Particles.Size.Random = rand.Intn(2) == 0

	// Randomize particles opacity
	config.Particles.Opacity.Value = 0.1 + rand.Float64()*0.9
	config.Particles.Opacity.Random = rand.Intn(2) == 0

	// Randomize particles movement
	config.Particles.Move.Speed = float64(rand.Intn(10) + 1)
	directions := []string{"none", "top", "top-right", "right", "bottom-right", "bottom", "bottom-left", "left", "top-left"}
	config.Particles.Move.Direction = directions[rand.Intn(len(directions))]
	config.Particles.Move.Random = rand.Intn(2) == 0
	config.Particles.Move.Straight = rand.Intn(2) == 0

	// Randomize line linking
	config.Particles.LineLinked.Enable = rand.Intn(2) == 0
	if config.Particles.LineLinked.Enable {
		config.Particles.LineLinked.Distance = float64(rand.Intn(300) + 100)
		config.Particles.LineLinked.Opacity = 0.1 + rand.Float64()*0.9
		config.Particles.LineLinked.Width = float64(rand.Intn(5) + 1)
	}

	// Randomize interactivity
	config.Interactivity.Events.OnHover.Enable = rand.Intn(2) == 0
	hoverModes := []string{"grab", "bubble", "repulse"}
	config.Interactivity.Events.OnHover.Mode = hoverModes[rand.Intn(len(hoverModes))]

	config.Interactivity.Events.OnClick.Enable = rand.Intn(2) == 0
	clickModes := []string{"push", "remove", "bubble", "repulse"}
	config.Interactivity.Events.OnClick.Mode = clickModes[rand.Intn(len(clickModes))]

	return config
}
