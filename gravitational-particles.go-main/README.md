# Particles-Go

A Go implementation of the popular [particles.js](https://github.com/VincentGarreau/particles.js) library, designed for integration with [Hugo](https://gohugo.io/) static websites.

## Features

- Generate particle configurations dynamically on the server using Go
- Multiple predefined particle presets (default, snow, night sky, bubbles, etc.)
- Pure gravity simulation with realistic particle movement 
- Customizable control panel with live updates
- Advanced interactivity options:
  - Standard hover modes (grab, repulse, bubble)
  - Click interactions (push, remove, repulse, bubble)
  - Particle attraction settings with adjustable strength
  - Custom cursor bounce effect (beta) - particles bounce off cursor like a screensaver
- Create random particle configurations
- Customizable through shortcode parameters
- Full compatibility with the original particles.js JavaScript library

## New Feature: Cursor Bounce (Beta)

We've implemented a new experimental feature that allows particles to bounce off the user's cursor like a classic screensaver:

- Enable via the "Enable cursor bounce" checkbox in the Interactivity section of the control panel
- Adjust bounce strength with the slider (controls the energy of the bounce)
- Adjust cursor radius to change the size of the bounce area
- Click "Apply Changes" to activate

**Note:** This feature is still in development and may not work consistently across all browsers and configurations. Currently, the feature:
- Works with the Pure Gravity simulation
- Provides visual feedback when particles bounce (brief flash)
- May experience occasional glitches with fast-moving particles

We're continuing to refine this feature to make it work more reliably and to integrate it more seamlessly with other interactivity options.

## Installation

1. First, ensure you have Go installed and set up on your system.

2. Clone this repository:
   ```bash
   git clone https://github.com/yourusername/particles-go.git
   cd particles-go
   ```

3. Build the project:
   ```bash
   go build
   ```

## Usage with Hugo

### Step 1: Add particles.js to your Hugo site

Download the original [particles.js](https://github.com/VincentGarreau/particles.js) library and add the `particles.min.js` file to your Hugo static directory:

```
your-hugo-site/
├── static/
│   └── js/
│       └── particles.min.js
```

### Step 2: Create the Hugo shortcode

Create a file at `layouts/shortcodes/particles.html` with the shortcode implementation provided in this repository.

### Step 3: Run the Go server

Run the included Go server alongside your Hugo site:

```bash
./particles-go
```

This will start a server on port 8080 that generates particle configurations.

### Step 4: Use the shortcode in your content

Now you can use the shortcode in your Hugo content files:

```markdown
<!-- Default configuration -->
{{< particles >}}

<!-- Use a preset -->
{{< particles preset="snow" >}}

<!-- Customize parameters -->
{{< particles color="#ff0000" number="150" size="5" >}}
```

## Available Presets

- `default`: Standard configuration with white particles and linking lines
- `snow`: Falling snow particles
- `nightsky`: A starry night sky effect with subtle twinkling
- `spacydots`: Connected dots that follow cursor movement
- `bubbles`: Floating bubble-like particles that react to mouse hover
- `gravity`: Particle system with gravitational forces applied between particles
- `strongGravity`: Enhanced gravity effect with a central "sun" object
- `pureGravity`: Custom physics-based gravity simulation separate from particles.js

## Shortcode Parameters

| Parameter | Description | Example |
|-----------|-------------|---------|
| id | Container element ID | `id="my-particles"` |
| preset | Use a predefined preset | `preset="snow"` |
| color | Particle color | `color="#ff0000"` |
| number | Number of particles | `number="150"` |
| shape | Particle shape type | `shape="circle"` |
| size | Particle size | `size="5"` |
| speed | Movement speed | `speed="3"` |
| direction | Movement direction | `direction="bottom"` |

## Control Panel

The demo includes an interactive control panel that allows real-time adjustment of particle properties:

### Particles Section
- **Number**: Adjust the quantity of particles displayed
- **Color**: Change particle color
- **Size**: Modify particle size
- **Opacity**: Control particle transparency
- **Shape**: Select from circle, triangle, star, or polygon

### Movement Section
- **Speed**: Adjust particle movement speed
- **Direction**: Set movement direction (none, top, bottom, left, right)
- **Particle Attraction**: Enable/disable gravitational-like attraction between particles
- **Attraction Strength**: Control the force of particle attraction

### Lines Section
- **Enable Lines**: Toggle connecting lines between particles
- **Line Color**: Set the color of connecting lines
- **Line Width**: Adjust the thickness of connecting lines
- **Connection Distance**: Set maximum distance for particles to connect

### Interactivity Section
- **On Hover**: Choose what happens when cursor hovers over particles (none, grab, bubble, repulse)
- **On Click**: Select click interaction behavior (none, push, remove, bubble, repulse)
- **Cursor Bounce**: Experimental feature for particles to bounce off cursor
  - **Bounce Strength**: Control the energy of particle bounce
  - **Cursor Radius**: Adjust the size of the cursor collision area

### Control Actions
- **Apply Changes**: Update the simulation with current settings
- **Reset to Default**: Restore original settings
- **Export Config**: Copy current configuration to clipboard

## Integration Options

### Option 1: Separate Services (Development)

Run the Hugo server and Go server as separate processes during development:

```bash
# Terminal 1
hugo server

# Terminal 2
./particles-go
```

### Option 2: Integrated Server (Production)

For production, you might want to:

1. Use a reverse proxy like Nginx to serve both your Hugo site and the Go API
2. Embed Hugo as a library in your Go application
3. Use Cloudflare Workers or similar to generate configurations client-side

## License

MIT License - See LICENSE file for details.

## Acknowledgements

- [particles.js](https://github.com/VincentGarreau/particles.js) - The original JavaScript library by Vincent Garreau
- [Hugo](https://gohugo.io/) - The static site generator

## Demo

To try out all features directly, you can run the included demo:

1. Open `demo.html` in your web browser
   ```bash
   # Windows
   start demo.html
   
   # macOS
   open demo.html
   
   # Linux
   xdg-open demo.html
   ```

2. Experiment with different presets using the preset buttons
3. Open the control panel in the top-right corner to customize settings
4. Try the cursor bounce feature:
   - Open the control panel
   - Check "Enable cursor bounce" in the Interactivity section
   - Adjust bounce strength and cursor size if desired
   - Click "Apply Changes"
   - Move your cursor around the screen to see particles bounce

You can also view a simpler demonstration in `simple-demo.html` for a more straightforward implementation.

## Known Issues

- The cursor bounce feature is still experimental and may not work consistently
- Particle attraction may sometimes cause particles to accelerate unexpectedly
- Very fast-moving particles may sometimes pass through the cursor without bouncing
- Some older browsers may experience performance issues with large numbers of particles
