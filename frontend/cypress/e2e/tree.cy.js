describe('Homepage', () => {
  beforeEach(() => {
    cy.visit("http://192.168.1.13:8080");
  })
  it('displays homepage', () => {
    cy.get('[data-test="welcome-header"]').contains("Welcome To Linktree App")
    cy.get('[data-test="tree-logo"]').should("be.visible")
    cy.contains("button", "Signup")
    cy.contains("button", "Login")
  })
  it('navigate to signup', () => {
    cy.get('[data-test="signup-btn"]').click()
    cy.get('[data-test="signup-header"]').contains("Sign up")
    cy.get('input').should('have.length', 4);
  })
  it('navigate to login', () => {
    cy.get('[data-test="login-btn"]').click()
    cy.get('[data-test="login-header"]').contains("Login")
    cy.get('input').should('have.length', 3)
  })
})
describe('Signup', () => {
  beforeEach(() => {
    cy.visit("http://192.168.1.13:8080");
    cy.get('[data-test="signup-btn"]').click()
    cy.get('[data-test="signup-header"]').contains("Sign up")
  })
  it('should allow user to signup', () => {
    cy.get('[data-test="username-input"]').type("testuser" + new Date().getTime())
    cy.get('[data-test="email-input"]').type("testuser@example.com")
    cy.get('[data-test="password-input"]').type("test1234")
    cy.get('[data-test="submit-btn"]').click()
    cy.url().should("eq", "http://192.168.1.13:8080/details")
    cy.get('[data-test="details-header"]').should("be.visible")
  })
  it('should not allow user to signup with no fields', () => {
    cy.get('[data-test="submit-btn"]').click()
    cy.on('window:alert', (text) => {
      expect(text).to.contains('Empty fields, please enter username, email and password')
    })
    cy.url().should("eq", "http://192.168.1.13:8080/signup")
  })
  it('should not allow user to signup due to weak password', () => {
    cy.get('[data-test="username-input"]').type("testuser")
    cy.get('[data-test="email-input"]').type("testuser@example.com")
    cy.get('[data-test="password-input"]').type("test")
    cy.get('[data-test="submit-btn"]').click()
    cy.on('window:alert', (text) => {
      expect(text).to.contains('Password must be at least 8 characters long')
    })
    cy.url().should("eq", "http://192.168.1.13:8080/signup")
  })
  it('should not allow same user to signup', () => {
    cy.get('[data-test="username-input"]').type("testuser")
    cy.get('[data-test="email-input"]').type("testuser@example.com")
    cy.get('[data-test="password-input"]').type("test1234")
    cy.get('[data-test="submit-btn"]').click()
    cy.on('window:alert', (text) => {
      expect(text).to.contains('User already exists, please login or try again')
    })
    cy.url().should("eq", "http://192.168.1.13:8080/signup")
  })

})
describe('Login', () => {
  beforeEach(() => {
    cy.visit("http://192.168.1.13:8080");
    cy.get('[data-test="login-btn"]').click()
    cy.get('[data-test="login-header"]').contains("Login")
  })
  it('should allow user to login in with correct credentials', () => {
    cy.get('[data-test="username-input"]').type("testuser")
    cy.get('[data-test="password-input"]').type("test1234")
    cy.get('[data-test="submit-btn"]').click()
    cy.on('window:alert', (text) => {
      expect(text).to.contains('Tree not found')
    })
    cy.url().should("eq", "http://192.168.1.13:8080/")
  })
  it('should not allow user to login in with empty fields', () => {
    cy.get('[data-test="submit-btn"]').click()
    cy.on('window:alert', (text) => {
      expect(text).to.contains('Empty fields, please enter username and password')
    })
  })
  it('should not allow user to login in with wrong credentials', () => {
    cy.get('[data-test="username-input"]').type("testwronguser")
    cy.get('[data-test="password-input"]').type("test1234")
    cy.get('[data-test="submit-btn"]').click()
    cy.on('window:alert', (text) => {
      expect(text).to.contains('Unauthorized user, please sign up or try again')
    })
  })

})

describe('Logout', () => {
  beforeEach(() => {
    cy.visit("http://192.168.1.13:8080");
    cy.get('[data-test="logout-btn"]').click()
  })
  it('should allow user to logout', () => {
    cy.on('window:confirm', (text) => {
      expect(text).to.contains('Are you sure you want to logout?')
      return true;
    })
    cy.on('window:alert', (text) => {
      expect(text).to.contains('Logout successful')
    })

  })
  it('should not allow user to logout if not logged in', () => {
    cy.on('window:confirm', (text) => {
      expect(text).to.contains('Are you sure you want to logout?')
      return true;
    })
    cy.on('window:alert', (text) => {
      expect(text).to.contains('You are not logged in')
    })

  })
  it('should not allow user to logout if cancel was clicker', () => {
    cy.on('window:confirm', (text) => {
      expect(text).to.contains('Are you sure you want to logout?')
      return false;
    })
    let alertShown = false;
    cy.on('window:alert', () => {
      alertShown = true;
    });

    cy.wrap(null).should(() => {
      expect(alertShown).to.be.false;
    });

  })


})