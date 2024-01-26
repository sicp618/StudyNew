import Image from "next/image";
import Counter from "./rigisterLogin";

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <div className="flex flex-col items-center justify-center">
        <Counter />
      </div>
    </main>
  );
}
