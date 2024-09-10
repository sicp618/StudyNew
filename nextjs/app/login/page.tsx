'use client';

import axios from "axios";
import { useState } from "react";
import { toast } from "react-toastify";

const loginUrl = `${process.env.NEXT_PUBLIC_API_HOST}/api/users/login`;

type User = {
  username: string;
  password: string;
};

export default function Login() {
  const [user, setUser] = useState<User>({
    username: "",
    password: "",
  });

  const loginAction = async () => {
    axios
      .post(loginUrl, { user })
      .then((res) => {
        setUser(res.data.user);
        toast.error(`登录成功`);
      })
      .catch((error) => {
        toast.error(`登录失败, ${error}`);
      })
  }

  return (
    <div>
      <div className="w-64">
        <input
          type="text"
          placeholder="username"
          value={user.username}
          onChange={(e) => setUser({ ...user, username: e.target.value })}
          className="input input-bordered w-full max-w-xs my-1"
        />
        <input
          type="password"
          placeholder="password"
          value={user.password}
          onChange={(e) => setUser({ ...user, password: e.target.value })}
          className="input input-bordered w-full max-w-xs my-1"
        />
        <button className="btn btn-primary w-full my-1" onClick={loginAction}>
          Login
        </button>
        <div className="form-control my-1">
          <label className="hover" >
            register
          </label>
        </div>
      </div>
    </div>
  );
}
