import { HttpException, HttpStatus, Injectable } from '@nestjs/common';
import { User } from './users.module';

@Injectable()
export class UsersService {
  private users: User[] = [];

  async register(user: User): Promise<User> {
    const userExists = this.users.find(u => u.username === user.username);
    if (userExists) {
      throw new HttpException('User already exists', HttpStatus.CONFLICT);
    }

    this.users.push(user);
    return user;
  }
}
