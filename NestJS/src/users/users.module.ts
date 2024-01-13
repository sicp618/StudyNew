import { Module } from '@nestjs/common';
import { UsersService } from './users.service';
import { UsersController } from './users.controller';
import { MongooseModule, Prop, SchemaFactory } from '@nestjs/mongoose';
import exp from 'constants';
import { User, UserSchema } from './user.schema';
import { SessionModule } from 'src/session/session.module';

@Module({
  imports: [
    MongooseModule.forFeature([{ name: User.name, schema: UserSchema }]),
    SessionModule,
  ],
  providers: [UsersService],
  controllers: [UsersController],
})
export class UsersModule {}
