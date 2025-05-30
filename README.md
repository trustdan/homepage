# Daniel's Homepage

A simple personal homepage built with HTML, CSS, and particles.js animations.

## Features

- Interactive particles background with gravitational effect
- Customizable particle settings via the control panel
- Responsive design that works on desktop and mobile

## Running Locally

There are multiple ways to run this project locally:

### Method 1: Open Directly

Simply open the HTML file directly in your browser:
```
gravitational-particles.go-main/index.html
```

You can use the included batch file:
```
.\open-homepage.bat
```

### Method 2: Using Node.js Server (Recommended)

1. Install dependencies:
```
npm install
```

2. Start the server:
```
npm start
```

3. Open your browser and go to:
```
http://localhost:3000/
```

You can also use the batch file:
```
.\serve-node.bat
```

### Method 3: Using Python's HTTP Server

1. Run the Python server:
```
cd gravitational-particles.go-main
python -m http.server
```

2. Open your browser and go to:
```
http://localhost:8000/
```

You can also use the PowerShell script:
```
.\serve.ps1
```

Or the batch file:
```
.\serve-python.bat
```

## Deployment

This project is ready to deploy to services like GitHub Pages or Netlify.

### Netlify Deployment

1. Push your repository to GitHub
2. Log in to Netlify and click "New site from Git"
3. Choose your repository
4. Configure:
   - Build command: (leave blank)
   - Publish directory: gravitational-particles.go-main
5. Click "Deploy site"

The included `netlify.toml` file is already configured for deployment, so Netlify will automatically use these settings.

## Customization

You can customize the particles effect by:

1. Clicking the "Particles.js Controls" in the top right corner
2. Adjusting the settings as desired
3. Clicking "Apply Changes"

## License

MIT 