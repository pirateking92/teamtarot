import { render, screen } from '@testing-library/svelte'
import userEvent from '@testing-library/user-event'
import { expect, test, vi } from 'vitest'
import { mockWindowCassandra } from '../test_utilities/mock-window';
import CassandraPage from './+page.svelte'

test('static greeting and button are displayed - cassandra page', async () => {
    render(CassandraPage)

    await new Promise((resolve) => setTimeout(resolve, 27500));

    const introText = screen.getByText(/Let me guide you through the land where fate/iu)
    const enterButton = screen.getByRole('button', { name: 'Enter the Parlor' })

    expect(enterButton).toBeInTheDocument()
    expect(introText).toBeInTheDocument()
}, 30000)

vi.stubGlobal('window', mockWindowCassandra);

test('route changes on button click - cassandra to card_display', async () => {
    const user = userEvent.setup()
    render(CassandraPage)

    await new Promise((resolve) => setTimeout(resolve, 27500));

    const introText = screen.getByText(/Let me guide you through the land where fate/iu)
    const enterButton = screen.getByRole('button', { name: 'Enter the Parlor' })
    const inputField = screen.getByPlaceholderText("Cassandra, I'd like to know about...")
    expect(enterButton).toBeInTheDocument()
    expect(introText).toBeInTheDocument()
    expect(inputField).toBeInTheDocument()

    expect(mockWindowCassandra.location.href).toBe('/cassandra?name=Paul');

    await user.type(inputField, "I am sad.");
    await user.click(enterButton);

    expect(mockWindowCassandra.location.href).toBe('/cards?name=&userstory=I%20am%20sad.');
}, 30000)