import express, { Request, Response, NextFunction, urlencoded, json } from 'express';
import { createServer } from 'http';
import { Server, Socket } from 'socket.io';
import cors from 'cors';
import { RabbitMQ } from './rabbitmq/client';

const rabbitMq = new RabbitMQ('amqp://khanh:handsome@localhost:5672', 'notification', null)
const rabbitMqTest = new RabbitMQ('amqp://khanh:handsome@localhost:5672', 'test', null)
const app = express();
app.use(cors());
app.use(json());
app.use(urlencoded({ extended: true }));
const server = createServer(app);
const port = 7789;
const io = new Server(server, {
  cors: {
    origin: '*',
  },
});

io.on('connection', (socket: Socket) => {
  console.log('a user connected', socket.id);
  io.emit('data_update', 'a new user join to us');
  socket.on('data_update', (args: any) => {
    console.log(args);
  })
  socket.on('disconnect', () => {
    console.log('a user disconnect');
  });
});

server.listen(port, () => {
  console.log(`Timezones by location application is running on port ${port}.`);
});
app.get('/', (req: Request, res: Response, next: NextFunction) => {
  res.status(200).json({ message: 'hi there!' });
});

app.post('/notifications', (req: Request, res: Response, next: NextFunction) => {
  const data = req.body;
  const success = io.emit('data_update', data)
  if (success) {
    res.status(200).json({ message: "your notification has sent to clients" });
  } else {
    res.status(500).json({ message: "fail to send your message to client" });
  }
})
