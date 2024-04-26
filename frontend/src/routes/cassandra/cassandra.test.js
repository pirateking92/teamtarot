import { render, screen } from '@testing-library/svelte'
import userEvent from '@testing-library/user-event'
import { expect, test, vi } from 'vitest'
import { mockWindowCassandra } from '../test_utilities/mock-window';
import CassandraPage from './+page.svelte'

test('static greeting and button are displayed - cassandra page', () => {
    render(CassandraPage)

    const introText = screen.getByText(/But be warned - the truths we may uncover could shake the very foundations of what you think you know./iu)
    const enterButton = screen.getByRole('button', { name: 'Enter' })

    expect(enterButton).toBeInTheDocument()
    expect(introText).toBeInTheDocument()
})

vi.stubGlobal('window', mockWindowCassandra);

test('route changes on button click - cassandra to card_display', async () => {
    const user = userEvent.setup()
    render(CassandraPage)

    const introText = screen.getByText(/But be warned - the truths we may uncover could shake the very foundations of what you think you know./iu)
    const enterButton = screen.getByRole('button', { name: 'Enter' })

    expect(enterButton).toBeInTheDocument()
    expect(introText).toBeInTheDocument()

    expect(mockWindowCassandra.location.href).toBe('/cassandra');

    await user.click(enterButton);

    expect(mockWindowCassandra.location.href).toBe('/card_display');
})