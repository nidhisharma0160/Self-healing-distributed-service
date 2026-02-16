from flask import Flask, jsonify
import threading
import time

app = Flask(__name__)
leak_storage = []
chaos_active = False

def memory_leak():
    global chaos_active
    while chaos_active:
        # Leak ~10MB every second
        leak_storage.append(" " * (10**7))
        time.sleep(1)

@app.route('/health')
def health():
    return jsonify(status="healthy"), 200

@app.route('/chaos', methods=['POST'])
def trigger_chaos():
    global chaos_active
    if not chaos_active:
        chaos_active = True
        threading.Thread(target=memory_leak, daemon=True).start()
        return jsonify(message="Chaos started! Watch the memory climb."), 200
    return jsonify(message="Chaos already running."), 400

if __name__ == '__main__':
    app.run(port=5000)