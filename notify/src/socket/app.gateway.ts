import { Logger } from "@nestjs/common";
import { MessageBody, OnGatewayConnection, OnGatewayDisconnect, OnGatewayInit, SubscribeMessage, WebSocketGateway, WebSocketServer } from "@nestjs/websockets";
import { Socket, Server } from 'socket.io';

@WebSocketGateway({ cors: true })
export class AppGateway implements OnGatewayInit, OnGatewayConnection, OnGatewayDisconnect {
    @WebSocketServer() server: Server;
    private logger: Logger = new Logger('AppGateway');

    afterInit(server: any) {
        this.logger.log('server init');
    }
    handleConnection(client: Socket, ...args: any[]) {
        this.logger.log('new connection: ', client.id, args);
    }
    handleDisconnect(client: Socket) {
        this.logger.log('disconnect: ', client.id);
    }

    @SubscribeMessage('message')
    handleMessage(@MessageBody() message: string): void {
        console.log('got new message');
        this.server.emit('message', message);
    }

}