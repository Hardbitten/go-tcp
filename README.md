
# Go-TCP 🌐

Hey there! 👋 Welcome to **Go-TCP** – a flexible Go-based TCP server and client framework perfect for building things like chat systems, data exchanges, and of course, basic game servers 🎮. Whether you're developing a multiplayer game or looking for a simple messaging system, this project has got you covered!

---

## 🚀 Why Go-TCP?

Here's a quick look at what makes Go-TCP pretty awesome:

- 📦 **Binary Protocol**: Speedy serialization and deserialization, keeping data light and fast.
- 🔄 **Opcode-Based Packets**: Modular, opcode-driven messages make managing different message types a breeze.
- ⚙️ **Concurrent Connections**: Handles multiple clients effortlessly with Goroutines.
- 📨 **Message Queueing**: Packet queues ensure messages stay organized and arrive in order.
- 🔗 **Connection Management**: Keeps things clean by disconnecting inactive clients.

---

## 🛠 Get Started!

Setting up Go-TCP is simple. Here’s how:

1. **Clone the repository**:

    ```bash
    git clone https://github.com/Hardbitten/go-tcp.git
    cd go-tcp
    ```

2. **Set up dependencies**:

    ```bash
    go mod tidy
    ```

3. **Build the project**:

    ```bash
    go build
    ```

---

## 🎨 Structure and Design

### 📑 Packet Structure

Every message is wrapped in a neat little packet with two main ingredients:

1. **Opcode**: Identifies the message type (like player movement or chat).
2. **Data Payload**: The actual info being sent, in binary format.

Here’s a quick sample of how packets are structured:

```go
func sendMessage(conn net.Conn, opcode uint16, data []byte) error {
    buffer := new(bytes.Buffer)
    binary.Write(buffer, binary.LittleEndian, opcode)
    buffer.Write(data)
    _, err := conn.Write(buffer.Bytes())
    return err
}
```

### 🔍 Serialization and Deserialization with ByteBuffer

**Go-TCP** includes a custom `ByteBuffer` package (`utils/ByteBuffer.go`) for serializing and deserializing data, making data processing efficient and compact.

- **Example: Handling Incoming Messages**  
  The server listens for specific opcodes and processes them based on incoming data from clients.

  ```go
  package handlers

  import (
      "main/models"
      op "main/opcodes"
      "main/utils"
  )

  // Define a type for handler functions
  type HandlerType func(data *utils.ByteBuffer, player *models.Player)

  var Handlers = map[uint16]HandlerType{
      op.CMSG_OPCODE_AUTH_LOGIN: HandleAuthLogin,
      op.CMSG_OPCODE_MATCH_FIND_START: HandleMatchFindStart,
      // More handlers here...
  }
  ```

- **Deserialization Example (Incoming Messages)**  
  Here’s how i handle player entry into the world using deserialization:

  ```go
  func HandlePlayerEnterWorld(data *utils.ByteBuffer, player *models.Player) {
      bf := serializers.SerializePlayerEnterWorld(player.ID, player.Position, player.Rotation, player.Name)
      player.BroadcastLobby(bf)
      fmt.Printf("Player [%d] joined the game.", player.ID)
  }
  ```

- **Serialization Example (Outgoing Messages)**  
  Here’s an example of converting data to binary to send to the client:

  ```go
  func SerializePlayerEnterWorld(PlayerID uint32, Position utils.Vector3, Rotation float32, Name string) *utils.ByteBuffer {
      buffer := utils.NewByteBuffer()
      buffer.WriteUInt16(opcodes.SMSG_OPCODE_PLAYER_ENTER_WORLD)
      buffer.WriteUInt32(PlayerID)
      buffer.WriteFloat(Position.X)
      buffer.WriteFloat(Position.Y)
      buffer.WriteFloat(Position.Z)
      buffer.WriteFloat(Rotation)
      buffer.WriteBytesWithLength([]byte(Name), 30) // Fixed 30-byte length
      return buffer
  }
  ```

---

## 🔄 Message Queueing & Client Connections

Each client gets its own message queue, which is great for keeping things smooth and orderly even with multiple messages flying around. Connections are managed with **Goroutines**, ensuring every client gets its own space to operate concurrently.

### 🤹 Multi-Client Connection Handling

Go-TCP cleans up as it goes by removing inactive connections, so the server doesn’t get bogged down by ghost clients.

---

## 🧪 Try It Out!

Let’s get a quick multiplayer scenario going:

1. **Start the Server**: Open a terminal and run `go run server.go`.
2. **Launch Clients**: Connect and listen to the TCP server using the provided IP and port.
3. **Send Messages**: Type away in each client to see your message system in action!

---

## ⚙️ Configuration

You can tweak things like the server IP, port, buffer sizes, and timeout settings right in `server.go` and `client.go` to match your project needs.

---

## 🌐 Use Cases

This project is like a Swiss Army knife for real-time communication! Here’s what it can be used for:

- 🎮 **Game Servers**: Create multiplayer fun with ease!
- 💬 **Chat Systems**: Real-time messaging without a hitch.
- 📊 **Data Streaming**: Great for IoT and live data feeds.
- 🔍 **Or whatever else you can dream up! 😄**

---

## 🤝 Contributing

Got ideas or improvements? i would love to see what you’ve got!

1. **Fork this repository** to make it yours.
2. **Create a feature branch** (`git checkout -b your-feature-branch`).
3. **Commit your awesome changes** with clear messages.
4. **Open a pull request** to share your work with the world.
---

Happy coding, and may Go-TCP make your real-time projects as fun and smooth as possible! 😊
