# Bidding System:

## _A bidding system, Machine Coding_ [Interview Problem]

- There will be sellers who can create auctions for items they want to
sell. They should be able to specify the lowest and highest bid that can
be placed along with a bid finalization time.
- There will be buyers who can place bids between lowest and highest
bid price. Buyers should be able to change bids till finalization time.
- Bids should conclude automatically on bid finalization time.
- On bid finalization the Bidding System should print on console the
winning bid with buyer and price.
- The bidding system should also print the sellers profit. For now
Seller's profit = winning bid’s price.

### Assumptions
- Sellers and Buyers can be assumed registered users.
- As a communication channel we are printing on console, later we can add other channels.
- Sellers profit calculation logic may change later.
- Bids can be stored In-Memory.

### Expectations:
- Code should be clean, modular, testable and extendable with minimal
changes. Follow SOLID principles.
- Code must compile and work.
- Handle edge cases

###Examples:
- Create Bid: Sell an antique coin. Lowest price: Rs. 100000, Highest Price: Rs. 10000000, Finalization time: 2021-06-16T21:00:00Z
- Bid1: 200000, Bid2: 5000000, Bid3: 5300000, Bid4: 600000
- At finalization time: Bid3 should be selected and Sellers profit should be **5300000**.
