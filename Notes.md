# Shortly - A URL shortenly service

## Pages

### Homepage 

    /

Add a link, customise the link


### Follow link

    /sha


Go to the URL, we should asynchronously store:

* Referer
* User-agent
* Timestamp
* One-way hash of some headers.

If the URL has expired, redirect to /app/expired
Send info in a 302

### Get info

    /app/info/<name>
    
Shows statistics about a given URL, based on /api/stats/ and eventually analytics.

### Expiry page

    /app/expired
 
Static page, explains expiry.


## API Actions

### Get URL (with/without name) [GET]

    /api/create

Creates a new shortcut, this is just a randomly generated
SHA which the user can customise.


### Store URL [POST]

    /api/store
    
Stores the shortcut with the given URL, along with timestamp (for expiring shortcuts).
    
### URL Info [GET]

    /api/info/<name>
    
Gets info about the specified name.
The URL that it redirects to mainly.

### URL Stats [GET]

    /api/stats/<name>
    
Gets statistics about the specified name. Access count for last 24 hours, week, month, all-time.


## Missing

### Auth.

- Authenicating users.
- Unauthenticated users have links that expire in 24 hours.

### Analytics

- What can we do with the info we collected?
- What visualisations can we provide?

### Naughty people

- Loop detection
- Validating URLs. Checking them
- Throttling _/api/create_ and _/api/store_


### Simple configuration

- How do we simplify configuration?

