// app/result/[requestID]/page.tsx
"use client";

import { useEffect, useState } from "react";

interface JSONCard {
  CardName: string;
  Type: string;
  MeaningUp: string;
  MeaningReverse: string;
  Description: string;
  ImageName: string;
  Reversed: boolean;
}

export default function ResultPage({
  params,
}: {
  params: { requestID: string };
}) {
  const [cards, setCards] = useState<JSONCard[]>([]);
  const [interpretation, setInterpretation] = useState<string>("");

  useEffect(() => {
    const fetchCards = async () => {
      const response = await fetch(`/cards/${params.requestID}`);
      if (response.ok) {
        const data = await response.json();
        setCards(data.cards);
        setInterpretation(data.interpretation);
      } else {
        console.error("Failed to load tarot reading.");
      }
    };

    fetchCards();
  }, [params.requestID]);

  return (
    <div>
      <h1>Your Tarot Reading</h1>
      <div>
        {cards.map((card, index) => (
          <div key={index}>
            <h2>{card.CardName}</h2>
            <p>{card.Description}</p>
            <p>{card.Reversed ? card.MeaningReverse : card.MeaningUp}</p>
          </div>
        ))}
      </div>
      <div>
        <h2>Interpretation</h2>
        <p>{interpretation}</p>
      </div>
    </div>
  );
}
