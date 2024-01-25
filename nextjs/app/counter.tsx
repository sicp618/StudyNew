"use client";

import axios from "axios";
import { useState } from "react";

const registerUrl = `${process.env.NEXT_PUBLIC_API_HOST}/api/register`;
const loginUrl = `${process.env.NEXT_PUBLIC_API_HOST}/api/login`;

type User = {
  username: string;
  password: string;
  email: string;
  id: number;
};

export default function Counter() {
  const [user, setUser] = useState<User | null>(null);

  const registerAction = async () => {
    const data = {
      user: {
        username: "user2",
        password: "123456",
        email: "user2@mail",
        id: 0,
      },
    };
    axios
      .post(registerUrl, data)
      .then((res) => {
        setUser(res.data.user);
      })
      .catch((error) => {
        console.log("error", error);
      });
  };

  const loginAction = async () => {
    console.log("loginAction", loginUrl, process.env.NEXT_PUBLIC_API_HOST, process.env);
    const loginData = {
      user: {
        username: "user2",
        password: "123456",
        email: "user2@mail",
        id: 0,
      },
    };
    axios
      .post(loginUrl, loginData)
      .then((res) => {
        console.log("res", res);
        setUser(res.data.user);
      })
      .catch((error) => {
        console.log("error", error);
      });
  };

  return (
    <div>
      <button onClick={registerAction}>register</button>
      <p />
      <button onClick={loginAction} hidden={!!user}>
        login
      </button>
      <p />
      <label>{user ? user.username : ""}</label>
    </div>
  );
}
