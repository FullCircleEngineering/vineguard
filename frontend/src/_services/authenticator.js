import { BehaviorSubject } from 'rxjs';

// the state of our currentUser dictates the highest-level behavior of our app.
// let it be a Subject that can be subscribed to to modulate behavior across the app.

const currentUserSubject = new BehaviorSubject(localStorage.getItem("cachedCurrentUser"));

const authenticator = {
    currentUser: currentUserSubject,
    setUserEmail,
    setUserToken,
    logout
};

function setUserToken(token) {
    localStorage.setItem("cachedCurrentUserToken", token)
}

function setUserEmail(email) {
    currentUserSubject.next(email);
    localStorage.setItem("cachedCurrentUser", email);
}

function logout() {
    currentUserSubject.next(null);
    localStorage.removeItem("cachedCurrentUser");
}

export { authenticator };
