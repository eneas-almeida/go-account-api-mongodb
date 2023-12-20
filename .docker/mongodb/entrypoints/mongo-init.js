db = db.getSiblingDB('customers');

db.createUser({
    user: 'customers',
    pwd: 'customers',
    roles: [{ role: 'readWrite', db: 'customers' }],
});
