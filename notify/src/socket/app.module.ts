import { Module } from "@nestjs/common";
import { AppGateway } from "./app.gateway";

@Module({
    providers: [AppGateway],
})
export class AppModule {}