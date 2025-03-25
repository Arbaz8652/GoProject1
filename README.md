# Golang WebSocket Chat

This is a simple **one-on-one chat application** built using **Golang WebSockets** and a basic web client. Users can connect, send, and receive messages in real-time.

## 🚀 Features
- Real-time **one-on-one messaging** using WebSockets.
- Supports multiple users.
- Simple web client for sending and receiving messages.
- Built using **Gorilla WebSocket**.

## 📌 Prerequisites
- Go **1.18+** installed
- Web browser (for testing the chat)

## 📦 Installation
1. Clone this repository:
   ```sh
   git clone https://github.com/Arbaz8652/GoProject1.git
   cd golang-websocket-chat
   ```

2. Install dependencies:
   ```sh
   go get -u github.com/gorilla/websocket
   ```

## 🏃 Running the Server
1. Start the WebSocket server:
   ```sh
   go run server.go
   ```
2. The server will start on **http://localhost:8080**.

## 💻 Running the Client
1. Open `index.html` in two different browsers or tabs.
2. Enter **different usernames** (e.g., `Alice` in one tab and `Bob` in another).
3. Use the **"To"** field to send messages to another user.

## 🔧 API Endpoints
### **1. WebSocket Connection**
- **Endpoint:** `/ws?username={username}`
- **Method:** `GET`
- **Description:** Connects a user to the WebSocket server.

### **2. Send Message**
- **Endpoint:** `/send?sender={sender}&receiver={receiver}&message={message}`
- **Method:** `GET`
- **Description:** Sends a message from `sender` to `receiver`.

## 🏗 Future Enhancements
- ✅ Store messages in a database (e.g., MongoDB, PostgreSQL)
- ✅ Add authentication (JWT-based login system)
- ✅ Improve the UI with chat history and better styling

## 📜 License
This project is **open-source** under the MIT License.

---
Made with ❤️ using Golang & WebSockets! 🚀

