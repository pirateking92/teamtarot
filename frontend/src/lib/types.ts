// Shared types mirroring the Go backend JSON shapes.

export interface Card {
	id: number;
	name: string;
	nameShort: string;
	arcana: 'major' | 'minor';
	suit: string;
	meaningUp: string;
	meaningReverse: string;
	description: string;
	imageFileName: string;
}

export interface ReadingCard {
	card: Card;
	positionIndex: number;
	positionLabel: string;
	reversed: boolean;
}

export type ReadingStatus = 'pending' | 'complete' | 'error';

export interface Reading {
	id: string;
	querentName: string;
	story: string;
	spreadType: string;
	status: ReadingStatus;
	interpretation: string;
	createdAt: string;
	cards: ReadingCard[];
}

export interface SpreadPosition {
	index: number;
	label: string;
}

export interface Spread {
	type: string;
	name: string;
	description: string;
	positions: SpreadPosition[];
}

export interface User {
	id: string;
	email: string;
	displayName: string;
	createdAt: string;
}

export interface DailyCard {
	date: string;
	card: Card;
	reversed: boolean;
}

/** Public URL for a card face image (served from static/cards). */
export function cardImage(card: Card): string {
	return `/cards/${card.imageFileName}`;
}
