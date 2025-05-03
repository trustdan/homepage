const express = require('express');
const path = require('path');
const app = express();
const port = process.env.PORT || 3000;

// Serve static files from the gravitational-particles.go-main directory
app.use(express.static(path.join(__dirname, 'gravitational-particles.go-main')));

// Redirect root to index.html
app.get('/', (req, res) => {
  res.redirect('/index.html');
});

// Start server
app.listen(port, () => {
  console.log(`Server running at http://localhost:${port}/`);
  console.log(`Open http://localhost:${port}/ to view your homepage`);
}); 