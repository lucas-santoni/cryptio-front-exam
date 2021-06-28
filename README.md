# Cryptio Front Exam

Hello! If you are reading this, it means that you are interested in taking a
**front-end position** at Cryptio.


## What is this about?

The objective of this test is to **validate your front-end skills**. You will be
given an **API exposing some data**, and your goal will be to **design and
implement a view** presenting this data to the end user.


## What is in the repository?

This repository is organized as follows:

- A Go back-end in the `back/` directory. You **do not** have to change the code
of this program. You don't even have to read it. You are only supposed to
consume the API it exposes. A Swagger documentation will help you explore this
API.
- A TypeScript front-end in the `front/` directory. It is obviously **not
complete**, we just have bootstrapped it for you. Actually, if you want, you can
completely **ignore it** (see [Technical
Considerations](#technical-considerations)).
- Screenshots from other products which might inspire you in the `gallery/`
directory.

## What am I supposed to do?

You are given an API, as well as an (hopefully!) complete Swagger documentation.
Please see [the API's README](back/README.md) for instructions on how to run
this API and access its associated Swagger. Your goal is to **consume the data**
exposed by this API, and **present it** to the end user.

The view you are going to work on is a **dashboard**. The general goal of this
view is to **display insights and metrics** about the portfolio of the end user.
The reason an user visits his dashboard is to answer questions such as:

- How much is my portfolio worth right now?
- What is the repartition of my portfolio? What assets do I own?
- How has my portfolio evolved over time?

The dashboard is organized into **widgets**. Each widget has a specific purpose
and answers at least one question related to the user's portfolio. Here are two
widgets you will implement:

### Portfolio Assets

This widget takes the form of **a table**. Each **row** corresponds to an
**asset**. The rows are **sorted by USD value**. The more of a particular asset
is worth in the portfolio, the higher it will be in the list. **For each row**,
you should display the following information:

- Asset logo
- Asset name
- Asset symbol
- Percentage of this asset in the portfolio's USD value
- Asset volume (how much of this asset is in the portfolio)
- Asset USD value (the current price of the volume)

An user's portfolio can contain a **very large amount of assets**, but the
objective of this widget is to **only show the top assets**. The widget should
only display **at most nine assets at a time**. The tenth row will be *the
rest*.

This widget is **not interactive**.

### Portfolio Balance

This widget takes the form of **a graph**. The x-axis is the time: each tick
represents a day. The y-axis is the value: each tick represents the total
portfolio value (in USD). Hence, each point of the graph is the portfolio USD
value at a certain day. You should display a total of two months on data, which
corresponds to around sixty points.

This widget is **interactive**: when the user's mouse hovers the graph, the
hovered date as well as the associated portfolio value will be clearly visible.
When the user does not interact with the widget, the **current portfolio
balance** should be clearly visible.

Make sure the widgets you design are responsive. They should look good on a
ultra-wide monitor, as well as on a phone. We will test your project on the
following devices:

- iPhone 5
- iPhone 12
- MacBook Pro 13 (2016)
- MacBook Pro 15 (2016)
- Random 1080p monitor
- Random 4k monitor


## Technical Considerations

You **must** use the following technologies:

- React
- TypeScript

You **may** use any flavor or React you want. We bootstrapped a Next.js
front-end for you in the `front/` directory of this repository but feel free to
ignore it and bootstrap your own project.

It is **mandatory** that you take advantage of TypeScript and maintain a type
safety.

You **are free to use any third-party dependency, except for *complete* React UI
frameworks such as Material UI or Chakra UI**. We want you to design your own
components. **CSS frameworks such as Bootstrap or Bulma are also forbidden**: we
want to read actual CSS code from you. However, CSS utilities like
Normalize.css are fine.

You **are free to use any HTTP client**, as long as you can maintain type
safety.


## Do you provide any mock-up?

We don't! Although we work with designers from time to time, **we usually don't
have mock-ups for the features we implement at Cryptio**. However, we took of few
screenshots of other products, which should be a great source of inspiration for
you. [Click here](./gallery/) to access the gallery.

**Note that this is definitely going to change. We aim at hiring a designer in
the near future.**


## Am I working for free here?

You are not! **None of your design/code/implementation produced for this test is
going to reach our production, ever.** We do have a dashboard page in the actual
Cryptio application, but we do not plan on updating it for now. It also fills
different purposes than the ones we are discussing in this test.


## How am I going to be evaluated?

You will be evaluated on the following topics:

- Proficiency with modern React (component life-cycle, hooks, state management,
data flow...)
- Your ability to properly fetch a remote API while maintaining correctness and
type safety
- Your ability to design responsive components
- Your ability to write clean, non-hacky CSS
- Your attention to detail, as well as you ability to put yourself in the user's
shoes
