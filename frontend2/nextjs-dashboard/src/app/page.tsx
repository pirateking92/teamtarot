"use client";

import { useState } from "react";
import axios from "axios";
import { useRouter } from "next/navigation";

const LandingPage = () => {
  const [name, setName] = useState("");
  const [query, setQuery] = useState("");
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      const response = await axios.post("/api/cards", { name, query });
      const { requestID } = response.data;
      router.push(`/result/${requestID}`);
    } catch (error) {
      console.error("Error submitting form:", error);
    }
  };

  return (
    <div>
      <h1>Tarot Card Reading</h1>
      <form onSubmit={handleSubmit}>
        <label>
          Name:
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
        </label>
        <label>
          Query:
          <textarea
            value={query}
            onChange={(e) => setQuery(e.target.value)}
          ></textarea>
        </label>
        <button type="submit">Get Reading</button>
      </form>
    </div>
  );
};

export default LandingPage;
