import { MiddlewareConsumer, Module, NestModule } from '@nestjs/common';
import { UsersModule } from './users/users.module';
import { MongooseModule } from '@nestjs/mongoose';
import { RedisModule } from './redis/redis.module';
import * as cookieParser from 'cookie-parser';
import { SessionModule } from './session/session.module';

@Module({
  imports: [
    MongooseModule.forRoot('mongodb://user:password@localhost:27017'),
    UsersModule,
    RedisModule,
    SessionModule,
  ],
})
export class AppModule implements NestModule {
  configure(consumer: MiddlewareConsumer) {
    consumer.apply(cookieParser()).forRoutes('*');
  }
}
