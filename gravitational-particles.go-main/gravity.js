// Custom gravity simulation with vanilla JavaScript
let gravityAnimationId = null;
let pureGravityActive = false;
// Track mouse position
let mouseX = null;
let mouseY = null;
let cursorRadius = 35; // Default cursor radius
let bounceFactor = 1.5; // Default bounce factor/strength
let debugMode = false; // Set to false to disable console logs
let lastBounceTime = 0; // Track last bounce for debugging

// Flag to track if bounce functionality is enabled
let bounceEnabled = false;

// Store handlers globally to allow removal
let currentMouseMoveHandler = null;
let currentMouseOutHandler = null;

// Function to update bounce settings
function updateBounceSettings(newBounceFactor, newCursorRadius) {
    bounceFactor = parseFloat(newBounceFactor) || 1.5;
    cursorRadius = parseInt(newCursorRadius) || 35;
    console.log(`Bounce settings updated: strength=${bounceFactor}, radius=${cursorRadius}px`);
}

// Expose functions to window
window.updateBounceSettings = updateBounceSettings;
window.applyBounceCursor = applyBounceCursor;

// Function to toggle debug mode with keyboard
function setupDebugToggle() {
    document.addEventListener('keydown', function(e) {
        // Press 'D' to toggle debug mode
        if (e.key === 'd' || e.key === 'D') {
            debugMode = !debugMode;
            console.log(`Debug mode ${debugMode ? 'enabled' : 'disabled'}`);
        }
        
        // Press '+' to increase cursor radius
        if (e.key === '+' || e.key === '=') {
            cursorRadius += 5;
            console.log(`Cursor radius increased to ${cursorRadius}px`);
        }
        
        // Press '-' to decrease cursor radius
        if (e.key === '-' || e.key === '_') {
            cursorRadius = Math.max(5, cursorRadius - 5);
            console.log(`Cursor radius decreased to ${cursorRadius}px`);
        }
    });
}

// NEW FUNCTION: Specifically for cursor bounce
function applyBounceCursor() {
    console.log('Starting gravity simulation WITH CURSOR BOUNCE');
    bounceEnabled = true;
    applyPureGravity();
}

function applyPureGravity() {
    // Check if we're applying with bounce or not
    const withBounce = bounceEnabled;
    console.log(`Applying pure gravity simulation ${withBounce ? 'WITH' : 'WITHOUT'} cursor bounce`);
    
    // Set up debug toggle
    setupDebugToggle();
    
    // First destroy existing particles if any
    if (typeof destroyParticles === 'function') {
        destroyParticles();
    }
    
    // Show debug instructions if bounce is enabled
    if (withBounce) {
        console.log("=== Cursor Bounce Debug Controls ===");
        console.log("Press 'D' to toggle debug mode");
        console.log("Press '+' to increase cursor radius");
        console.log("Press '-' to decrease cursor radius");
        console.log("===================================");
    }
    
    // Hide particles.js container
    const particlesContainer = document.getElementById('particles-js');
    particlesContainer.style.display = 'none';
    
    // Create our own canvas for gravity simulation
    let gravityCanvas = document.getElementById('gravity-canvas');
    if (!gravityCanvas) {
        gravityCanvas = document.createElement('canvas');
        gravityCanvas.id = 'gravity-canvas';
        gravityCanvas.style.position = 'fixed';
        gravityCanvas.style.top = '0';
        gravityCanvas.style.left = '0';
        gravityCanvas.style.width = '100%';
        gravityCanvas.style.height = '100%';
        gravityCanvas.style.zIndex = '1';
        gravityCanvas.style.backgroundColor = '#0a192f';
        document.body.appendChild(gravityCanvas);
    } else {
        gravityCanvas.style.display = 'block';
    }
    
    // Setup canvas and context
    const ctx = gravityCanvas.getContext('2d');
    gravityCanvas.width = window.innerWidth;
    gravityCanvas.height = window.innerHeight;
    
    // Only set up mouse tracking if bounce is enabled
    if (withBounce) {
        // Remove previous listeners if any exist
        if (currentMouseMoveHandler) {
            gravityCanvas.removeEventListener('mousemove', currentMouseMoveHandler);
        }
        if (currentMouseOutHandler) {
            gravityCanvas.removeEventListener('mouseout', currentMouseOutHandler);
        }
        
        // Define mouse event handlers
        currentMouseMoveHandler = function(e) {
            mouseX = e.clientX;
            mouseY = e.clientY;
            if (debugMode) {
                if (Math.random() < 0.01) {
                    console.log(`Mouse position: ${mouseX}, ${mouseY}`);
                }
            }
        };
        currentMouseOutHandler = function() {
            mouseX = null;
            mouseY = null;
            if (debugMode) console.log('Mouse left canvas');
        };
        
        // Add event listeners
        gravityCanvas.addEventListener('mousemove', currentMouseMoveHandler);
        gravityCanvas.addEventListener('mouseout', currentMouseOutHandler);
        
        console.log('🔴 Mouse tracking enabled - move your cursor over the canvas!');
    } else {
        // Reset mouse position variables if bounce is not enabled
        mouseX = null;
        mouseY = null;
    }
    
    // Update status
    const statusEl = document.getElementById('status');
    statusEl.textContent = withBounce 
        ? 'Current preset: Pure Gravity with Cursor Bounce' 
        : 'Current preset: Pure Gravity';
    
    // Define central sun
    const sun = {
        x: gravityCanvas.width / 2,
        y: gravityCanvas.height / 2,
        radius: 30,
        mass: 2000,
        color: '#ffdd00'
    };
    
    // Create planets
    const planets = [];
    const colors = ['#ff7e7e', '#7eff8e', '#7ee0ff', '#ffffff'];
    
    for (let i = 0; i < 150; i++) {
        const angle = Math.random() * Math.PI * 2;
        const distance = 100 + Math.random() * 150;
        const x = sun.x + Math.cos(angle) * distance;
        const y = sun.y + Math.sin(angle) * distance;
        
        const speed = Math.sqrt(sun.mass / distance) * 0.5;
        
        planets.push({
            x: x,
            y: y,
            radius: 2 + Math.random() * 4,
            mass: 1,
            color: colors[Math.floor(Math.random() * colors.length)],
            vx: Math.sin(angle) * speed,
            vy: -Math.cos(angle) * speed,
            flashTime: 0
        });
    }
    
    // Flag to track if simulation is active
    pureGravityActive = true;
    
    if (withBounce) {
        console.log('Simulation initialized with cursor bounce functionality enabled');
        console.log(`Move your cursor over the canvas (cursor radius: ${cursorRadius}px)`);
    } else {
        console.log('Pure gravity simulation initialized (no cursor bounce)');
    }
    
    // Animation loop
    function animate() {
        if (!pureGravityActive) {
            cancelAnimationFrame(gravityAnimationId);
            return;
        }
        
        // Clear canvas
        ctx.clearRect(0, 0, gravityCanvas.width, gravityCanvas.height);
        
        // Draw sun
        ctx.beginPath();
        ctx.arc(sun.x, sun.y, sun.radius, 0, Math.PI * 2);
        ctx.fillStyle = sun.color;
        ctx.fill();
        
        // Draw cursor area if mouse is on canvas and bounce is enabled
        if (withBounce && mouseX !== null && mouseY !== null) {
            ctx.beginPath();
            ctx.arc(mouseX, mouseY, cursorRadius, 0, Math.PI * 2);
            ctx.fillStyle = 'rgba(255, 255, 255, 0.3)'; // More visible
            ctx.strokeStyle = 'rgba(255, 255, 255, 0.8)';
            ctx.lineWidth = 2;
            ctx.fill();
            ctx.stroke();
            
            // Draw a clear center point
            ctx.beginPath();
            ctx.arc(mouseX, mouseY, 3, 0, Math.PI * 2);
            ctx.fillStyle = '#ffffff';
            ctx.fill();
        }
        
        // Count bounces in this frame for debugging
        let bouncesThisFrame = 0;
        
        // Update and draw planets
        for (const planet of planets) {
            // Calculate gravitational force
            const dx = sun.x - planet.x;
            const dy = sun.y - planet.y;
            const distSq = dx * dx + dy * dy;
            const dist = Math.sqrt(distSq);
            
            // Skip if too close to prevent extreme forces
            if (dist > sun.radius + planet.radius) {
                const force = sun.mass / distSq * 0.1;
                
                // Apply acceleration
                planet.vx += (dx / dist) * force;
                planet.vy += (dy / dist) * force;
            }
            
            // Check for collision with cursor (only if bounce is enabled)
            if (withBounce && mouseX !== null && mouseY !== null) {
                const dxCursor = mouseX - planet.x;
                const dyCursor = mouseY - planet.y;
                const distCursorSq = dxCursor * dxCursor + dyCursor * dyCursor;
                const distCursor = Math.sqrt(distCursorSq);
                
                // Calculate the path of the particle to detect fast-moving particles
                // that might skip over the cursor in a single frame
                const nextX = planet.x + planet.vx;
                const nextY = planet.y + planet.vy;
                const nextDxCursor = mouseX - nextX;
                const nextDyCursor = mouseY - nextY;
                const nextDistCursorSq = nextDxCursor * nextDxCursor + nextDyCursor * nextDyCursor;
                const nextDistCursor = Math.sqrt(nextDistCursorSq);
                
                // Detect if particle is currently colliding or will collide next frame
                const willCollide = distCursor < cursorRadius + planet.radius || 
                                   nextDistCursor < cursorRadius + planet.radius;
                
                // If planet is colliding with cursor
                if (willCollide) {
                    // Calculate bounce direction (normal to collision)
                    const nx = dxCursor / distCursor;
                    const ny = dyCursor / distCursor;
                    
                    // Calculate dot product of velocity and normal
                    const dot = planet.vx * nx + planet.vy * ny;
                    
                    // Only bounce if planet is moving toward the cursor
                    if (dot < 0) {
                        // Reflect velocity vector using normal
                        planet.vx = planet.vx - 2 * dot * nx;
                        planet.vy = planet.vy - 2 * dot * ny;
                        
                        // Add more energy to the bounce for visibility
                        planet.vx *= bounceFactor;
                        planet.vy *= bounceFactor;
                        
                        // Move planet outside of cursor to prevent sticking
                        const moveDistance = cursorRadius + planet.radius - distCursor + 2; // Add extra distance
                        planet.x = planet.x - nx * moveDistance;
                        planet.y = planet.y - ny * moveDistance;
                        
                        // Visual feedback - flash the particle brighter
                        planet.flashTime = 10; // Frames to show flash effect
                        
                        // Debug
                        bouncesThisFrame++;
                        lastBounceTime = Date.now();
                    }
                }
            }
            
            // Apply slight damping
            planet.vx *= 0.999;
            planet.vy *= 0.999;
            
            // Update position
            planet.x += planet.vx;
            planet.y += planet.vy;
            
            // Bounce off edges
            if (planet.x < 0 || planet.x > gravityCanvas.width) {
                planet.vx *= -0.9;
            }
            if (planet.y < 0 || planet.y > gravityCanvas.height) {
                planet.vy *= -0.9;
            }
            
            // Draw planet
            ctx.beginPath();
            ctx.arc(planet.x, planet.y, planet.radius, 0, Math.PI * 2);
            
            // Apply flash effect if recently bounced
            if (planet.flashTime > 0) {
                ctx.fillStyle = '#ffffff'; // White flash
                planet.flashTime--;
            } else {
                ctx.fillStyle = planet.color;
            }
            
            ctx.fill();
            
            // Draw line to sun
            const lineMaxDist = 250;
            if (dist < lineMaxDist) {
                ctx.beginPath();
                ctx.moveTo(planet.x, planet.y);
                ctx.lineTo(sun.x, sun.y);
                ctx.strokeStyle = `rgba(255, 255, 255, ${1 - dist / lineMaxDist})`;
                ctx.lineWidth = 1;
                ctx.stroke();
            }
        }
        
        // Log bounces for debugging
        if (withBounce && debugMode && bouncesThisFrame > 0) {
            console.log(`Bounces this frame: ${bouncesThisFrame}`);
        }
        
        // Request next frame
        gravityAnimationId = requestAnimationFrame(animate);
    }
    
    // Start animation
    animate();
    
    console.log(`Pure gravity simulation ${withBounce ? 'with cursor bounce' : ''} started`);
}

// Handle window resizing
window.addEventListener('resize', function() {
    const gravityCanvas = document.getElementById('gravity-canvas');
    if (gravityCanvas && pureGravityActive) {
        gravityCanvas.width = window.innerWidth;
        gravityCanvas.height = window.innerHeight;
    }
});

// Function to stop gravity simulation
function stopGravitySimulation() {
    if (pureGravityActive) {
        pureGravityActive = false;
        cancelAnimationFrame(gravityAnimationId);
        
        const gravityCanvas = document.getElementById('gravity-canvas');
        if (gravityCanvas) {
            gravityCanvas.style.display = 'none';
            
            // Remove event listeners using the stored handlers
            if (currentMouseMoveHandler) {
                gravityCanvas.removeEventListener('mousemove', currentMouseMoveHandler);
                currentMouseMoveHandler = null; // Clear stored handler
            }
            if (currentMouseOutHandler) {
                gravityCanvas.removeEventListener('mouseout', currentMouseOutHandler);
                currentMouseOutHandler = null; // Clear stored handler
            }
            console.log('Cleaned up gravity canvas event listeners.');
        }
        
        const particlesContainer = document.getElementById('particles-js');
        particlesContainer.style.display = 'block';
        
        // Reset bounce flag
        bounceEnabled = false;
        
        console.log('Stopped gravity simulation');
    } else {
        // Ensure particles-js is visible even if gravity wasn't active
        const particlesContainer = document.getElementById('particles-js');
        if (particlesContainer) particlesContainer.style.display = 'block';
        // Also ensure gravity canvas is hidden if it exists
        const gravityCanvas = document.getElementById('gravity-canvas');
        if (gravityCanvas) gravityCanvas.style.display = 'none';
    }
}

// Set up override of destroyParticles when the page loads
window.addEventListener('DOMContentLoaded', function() {
    if (typeof destroyParticles === 'function') {
        const originalDestroyParticles = destroyParticles;
        window.destroyParticles = function() {
            originalDestroyParticles();
            stopGravitySimulation();
        };
    }
}); 