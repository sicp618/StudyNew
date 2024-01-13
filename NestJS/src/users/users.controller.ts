import {
  Body,
  Controller,
  HttpException,
  HttpStatus,
  Post,
  Req,
  Res,
} from '@nestjs/common';
import { UsersService } from './users.service';
import { User } from './user.schema';
import { Response } from 'express';
import { SessionService } from '../session/session.service';

@Controller('api/users')
export class UsersController {
  constructor(private usersService: UsersService, private sessionService: SessionService) {}

  @Post('register')
  async register(@Body('user') user: User): Promise<{ user: User }> {
    try {
      console.log('register input:', user);
      const registeredUser = await this.usersService.register(user);
      return { user: registeredUser };
    } catch (error) {
      if (error instanceof HttpException) {
        throw error;
      }
      throw new HttpException(error.message, HttpStatus.INTERNAL_SERVER_ERROR);
    }
  }

  @Post('login')
  async login(
    @Res({ passthrough: true }) res: Response,
    @Body('user') user: User
  ): Promise<{ user: User }> {
    try {
      console.log('login input:', user);
      const loggedInUser = await this.usersService.login(user);

      let sessionId = await this.sessionService.createSession(user.username);
      res.cookie('nid', sessionId, {
        httpOnly: true,
        // secure: true,
        maxAge: 1000 * 60 * 60 * 24 * 7,
      });
      return { user: loggedInUser };
    } catch (error) {
      if (error instanceof HttpException) {
        throw error;
      }
      throw new HttpException(error.message, HttpStatus.INTERNAL_SERVER_ERROR);
    }
  }
}
