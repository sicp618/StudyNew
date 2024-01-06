import { UsersService } from './users.service';
import { HttpException, HttpStatus } from '@nestjs/common';

describe('UsersService', () => {
  let service: UsersService;

  beforeEach(() => {
    service = new UsersService();
  });

  it('should register a new user', async () => {
    const user = {
      username: 'test',
      password: 'test',
      id: 0,
      email: 'test@test',
    };
    const result = await service.register(user);
    expect(result).toEqual(user);
  });

  it('should throw an exception if the user already exists', async () => {
    const user = {
      username: 'test',
      password: 'test',
      id: 0,
      email: 'test@test',
    };
    await service.register(user);

    try {
      await service.register(user);
    } catch (e) {
      expect(e).toBeInstanceOf(HttpException);
      expect(e.getStatus()).toEqual(HttpStatus.CONFLICT);
      expect(e.getResponse()).toEqual('User already exists');
    }
  });
});
