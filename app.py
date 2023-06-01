import os
import time
from flask import Flask, render_template

app = Flask(__name__)

@app.route('/')
def animate_smiley():
	return render_template('/app.html')

if __name__ == "__main__":
	port = int(os.environ.get('PORT', 5000))
	app.run(debug=True, host='0.0.0.0', port=port)

