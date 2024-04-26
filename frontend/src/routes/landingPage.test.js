import { render, screen } from '@testing-library/svelte'
import userEvent from '@testing-library/user-event'
import { expect, test, vi } from 'vitest'
import { mockWindowLanding } from './test_utilities/mock-window';
import LandingPage from './+page.svelte'

test('static greeting and button are displayed - landing page', () => {
    render(LandingPage)

    const button = screen.getByRole('button', { name: 'Follow Cassandra' })
    const greetingText = screen.queryByText(/Ah, I've been expecting you, seeker. My name is Cassandra./iu)

    expect(button).toBeInTheDocument()
    expect(greetingText).toBeInTheDocument()
})

vi.stubGlobal('window', mockWindowLanding);

test('route changes on button click - landing to cassandra', async () => {
    const user = userEvent.setup()
    render(LandingPage)

    const button = screen.getByRole('button', { name: 'Follow Cassandra' })
    const greetingText = screen.queryByText(/Ah, I've been expecting you, seeker. My name is Cassandra./iu)

    expect(button).toBeInTheDocument()
    expect(greetingText).toBeInTheDocument()

    expect(mockWindowLanding.location.href).toBe('/');

    await user.click(button);

    expect(mockWindowLanding.location.href).toBe('/cassandra');
})