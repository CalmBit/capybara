<img src="capybara.svg" width="128px" height="128px" align="right">

# Capybara
Capybara is a [Mastodon](https://github.com/tootsuite/mastodon) clone - it will function practically indentically to
Mastodon, both under the hood and through its exposed APIs.

## Why?
For fun, and because although I love Eugen's work, I'm wondering if putting the backend in Go might make things
go faster.

## How?
All of the code will be designed independently of Mastodon, but the underlying schema and interface (including the
database layout and API route) will be, for compatibility and cleanliness purposes, identical to the best of my ability.

## Development/Deployment
Currently, this project is just barely getting started, so
I have no advice or recommendations aside from "this isn't
even close to working."

## License
Since Capybara is derivative, this will include two license notices - one for Capybara's side of the code (the 
application itself) and one for the API/database design (from Mastodon). Most of the credit should go to Eugen and
the beautiful people working on Mastodon - this is just a tribute.

---
### Capybara
Copyight (C) 2018 CalmBit.

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General 
Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) 
any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied 
warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more 
details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, 
see https://www.gnu.org/licenses/.

---
### Mastodon


Copyright (C) 2016-2018 Eugen Rochko & other Mastodon contributors (see AUTHORS.md)

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see https://www.gnu.org/licenses/.

---
