import os
import time
from flask import Flask, render_template

app = Flask(__name__)

@app.route('/')
def animate_smiley():
	return render_template('/app.html')
