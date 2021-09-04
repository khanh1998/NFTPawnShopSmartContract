import { Controller, Get } from '@nestjs/common';
import { AppService } from './app.service';
import { GrpcMethod } from '@nestjs/microservices';
import { app } from 'app';
import { Metadata } from '@grpc/grpc-js';
import { Observable, Subject } from 'rxjs';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @GrpcMethod('PushNotification', 'SendNotification')
  sendNotification(data: app.NotificationData, metadata?: Metadata): app.NotificationResult {
    console.log(data);
    const res: app.NotificationResult = {
      success: true,
    }
    return res;
  }
}
