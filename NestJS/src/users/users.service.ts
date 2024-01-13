import { HttpException, HttpStatus, Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { User, UserDocument } from './user.schema';

@Injectable()
export class UsersService {
  constructor(@InjectModel(User.name) private userModel: Model<UserDocument>) {}

  async register(user: User): Promise<User> {
    const userExists = await this.userModel.findOne({
      username: user.username,
    });
    if (userExists) {
      throw new HttpException('User already exists', HttpStatus.CONFLICT);
    }

    const createdUser = await this.userModel.create(user);
    return createdUser.save();
  }

  async login(user: User): Promise<User> {
    const userExists = await this.userModel.findOne({
      username: user.username,
    });
    if (!userExists) {
      throw new HttpException('User does not exist', HttpStatus.NOT_FOUND);
    }

    const passwordMatches = userExists.password === user.password;
    if (!passwordMatches) {
      throw new HttpException('Password is incorrect', HttpStatus.UNAUTHORIZED);
    }
    return userExists;
  }
}
