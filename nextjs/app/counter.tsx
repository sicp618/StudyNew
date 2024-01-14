"use client";

import { useState } from "react";

const registerAction = async () => {
  const data = {
    user: {
      username: "user2",
      password: "123456",
      email: "user2@mail",
      id: 0,
    },
  };
  const body = JSON.stringify(data);
  try {
    const res = await fetch("http://localhost:3000/api/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: body,
    });

    if (!res.ok) {
      throw new Error(res.statusText);
    }

    const result = await res.json();
    console.log("res", result);
  } catch (error) {
    console.log("error", error);
  }
};

export default function Counter() {
  const [count, setCount] = useState(0);

  return (
    <div>
      <p>You clicked {count} times</p>
      <button onClick={registerAction}>Click me</button>
    </div>
  );
}
