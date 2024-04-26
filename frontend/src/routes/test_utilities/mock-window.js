import { vi } from 'vitest';

export const mockWindowLanding = {
    location: {
        href: '/',
        assign: vi.fn(),
        replace: vi.fn(),
    },
};

export const mockWindowCassandra = {
    location: {
        href: '/cassandra',
        assign: vi.fn(),
        replace: vi.fn(),
    },
};

export const mockWindowCardDisplay = {
    location: {
        href: '/card_display',
        assign: vi.fn(),
        replace: vi.fn(),
    },
};