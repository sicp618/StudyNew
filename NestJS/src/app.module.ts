import { MiddlewareConsumer, Module, NestModule } from '@nestjs/common';
import { MongooseModule } from '@nestjs/mongoose';
import { RedisModule } from './redis/redis.module';
import { ConfigModule, ConfigService } from '@nestjs/config';
import * as cookieParser from 'cookie-parser';
import { SessionModule } from './session/session.module';
import { PrismaModule } from './prisma/prisma.module';
import { PostsService } from './posts/posts.service';
import { CommentsService } from './comments/comments.service';
import { LabelsService } from './labels/labels.service';
import { PostsController } from './posts/posts.controller';
import { CommentsController } from './comments/comments.controller';
import { LabelsController } from './labels/labels.controller';
import { PrismaService } from './prisma/prisma.service';
import { UsersService } from './users/users.service';
import { UsersController } from './users/users.controller';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
      envFilePath: ['.env.local', '.env'],
    }),
    MongooseModule.forRootAsync({
      imports: [ConfigModule],
      useFactory: async (configService: ConfigService) => ({
        uri: configService.get<string>('MONGODB_URI'),
      }),
      inject: [ConfigService],
    }),
    PrismaModule,
    RedisModule,
    SessionModule,
  ],
  providers: [PostsService, CommentsService, LabelsService, UsersService, PrismaService],
  controllers: [PostsController, CommentsController, LabelsController, UsersController],
})
export class AppModule implements NestModule {
  configure(consumer: MiddlewareConsumer) {
    consumer.apply(cookieParser()).forRoutes('*');
  }
}