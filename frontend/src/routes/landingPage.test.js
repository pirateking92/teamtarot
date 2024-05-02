import { render, screen } from '@testing-library/svelte'
import userEvent from '@testing-library/user-event'
import { expect, test, vi } from 'vitest'
import { mockWindowLanding } from './test_utilities/mock-window';
import LandingPage from './+page.svelte'

test('static greeting and button are displayed - landing page', async () => {
    render(LandingPage)

    await new Promise((resolve) => setTimeout(resolve, 11500));

    const button = screen.getByRole('button', { name: 'Follow Cassandra' })
    const greetingText = screen.queryByText(/Ah, I've been expecting you, seeker./iu)

    expect(button).toBeInTheDocument()
    expect(greetingText).toBeInTheDocument()
}, 20000)

vi.stubGlobal('window', mockWindowLanding);

test('route changes on button click - landing to cassandra', async () => {
    const user = userEvent.setup()
    render(LandingPage)

    await new Promise((resolve) => setTimeout(resolve, 11500));

    const button = screen.getByRole('button', { name: 'Follow Cassandra' })
    const greetingText = screen.queryByText(/Ah, I've been expecting you, seeker./iu)
    const inputField = screen.getByPlaceholderText('Enter your name')

    expect(button).toBeInTheDocument()
    expect(greetingText).toBeInTheDocument()
    expect(inputField).toBeInTheDocument()


    expect(mockWindowLanding.location.href).toBe('/');

    await user.type(inputField, "Paul");
    await user.click(button);

    expect(mockWindowLanding.location.href).toBe('/cassandra?name=Paul');
}, 20000)