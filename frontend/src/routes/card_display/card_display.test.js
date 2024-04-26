import { render, screen } from '@testing-library/svelte'
import userEvent from '@testing-library/user-event'
import { expect, test, vi } from 'vitest'
import { mockWindowCardDisplay } from '../test_utilities/mock-window';
import CardDisplayPage from './+page.svelte'

test('static text and button are displayed - card_display page', () => {
    render(CardDisplayPage)

    const headerText = screen.getByText(/Welcome to Cassandra's Parlor/iu)
    const leaveButton = screen.getByRole('button', { name: 'Leave the Parlor' })

    expect(headerText).toBeInTheDocument()
    expect(leaveButton).toBeInTheDocument()
})

vi.stubGlobal('window', mockWindowCardDisplay);

test('route changes on button click - cassandra to card_display', async () => {
    const user = userEvent.setup()
    render(CardDisplayPage)

    const headerText = screen.getByText(/Welcome to Cassandra's Parlor/iu)
    const leaveButton = screen.getByRole('button', { name: 'Leave the Parlor' })

    expect(headerText).toBeInTheDocument()
    expect(leaveButton).toBeInTheDocument()

    expect(mockWindowCardDisplay.location.href).toBe('/card_display');

    await user.click(leaveButton);

    expect(mockWindowCardDisplay.location.href).toBe('/');
})


test('dynamic cards are displayed - card_display page', async () => {
    render(CardDisplayPage)

    await new Promise((resolve) => setTimeout(resolve, 2000));

    const pastCards = screen.getAllByText(/Past/iu)
    const presentCards = screen.getAllByText(/Present/iu)
    const futureCards = screen.getAllByText(/Future/iu)

    const minorTypes = screen.queryAllByText(/Minor/iu);
    const majorTypes = screen.queryAllByText(/Major/iu);

    expect(pastCards.length).toBeGreaterThan(0);
    expect(presentCards.length).toBeGreaterThan(0);
    expect(futureCards.length).toBeGreaterThan(0);
    expect(minorTypes.length > 0 || majorTypes.length > 0).toBe(true);
})