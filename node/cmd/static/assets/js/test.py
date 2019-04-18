from flask import Flask, request
import json

app = Flask(__name__)



@app.route('/getSomething', methods=['GET'])
def getRequest():
    details = [{
	"name": "Apple-pie",
	"image_url": "https://www.simplyrecipes.com/wp-content/uploads/2014/09/apple-pie-vertical-b-1600-600x889.jpg",
	"details": "Tasty Granma Pie"
    }]
    return json.dumps(details);
"""
@app.route('/testRequest')
def requestTesting():
    return "I am working fine"
"""

if __name__ == '__main__':
    app.run(host= '127.0.0.1',port= '8080')
