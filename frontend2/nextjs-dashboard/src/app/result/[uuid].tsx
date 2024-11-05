"use client";

import { useState, useEffect } from "react";
import { useParams } from "next/navigation";
import axios from "axios";

const ResultPage = () => {
  const { uuid } = useParams();
  const [interpretation, setInterpretation] = useState("");
  const [cards, setCards] = useState<[]>([]);

  useEffect(() => {
    const fetchInterpretation = async () => {
      try {
        const response = await axios.get(
          `http://localhost:8082/api/interpretation/${uuid}`
        );
        setInterpretation(response.data.interpretation);
      } catch (error) {
        console.error("Error fetching interpretation:", error);
      }
    };

    const fetchCards = async () => {
      try {
        const response = await axios.get(
          `http://localhost:8082/api/cards/${uuid}`
        );
        setCards(response.data.cards);
      } catch (error) {
        console.error("Error fetching cards:", error);
      }
    };

    fetchInterpretation();
    fetchCards();
  }, [uuid]);

  return (
    <div>
      <h1>Tarot Card Reading</h1>
      <div>
        <h2>Interpretation:</h2>
        <p>{interpretation}</p>
      </div>
      <div>
        <h2>Cards:</h2>
        <ul>
          {cards.map((card, index) => (
            <li key={index}>
              {card.CardName}
              {card.Reversed && " (Reversed)"}
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
};

export default ResultPage;
