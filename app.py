import os
import time
from flask import Flask, render_template
from flask_cors import CORS # CORS plugin

# enable CORS
app = Flask(__name__)
CORS(app)

@app.route('/')
def animate_smiley():
	return render_template('/app.html')

if __name__ == "__main__":
	port = int(os.environ.get('PORT', 5000))
	app.run(debug=False, host='0.0.0.0', port=port)

