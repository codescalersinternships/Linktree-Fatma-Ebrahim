describe('Homepage', () => {
  beforeEach(()=>{
    cy.visit("http://192.168.1.13:8080"); 
  })
  it('displays homepage', () => {
    cy.get('[data-test="welcome-header"]').contains("Welcome To Linktree App")
    cy.get('[data-test="tree-logo"]').should("be.visible")
    cy.contains("button","Signup")
    cy.contains("button","Login")
  })
  it('navigate to signup',()=>{
    cy.get('[data-test="signup-btn"]').click()
    cy.get('[data-test="signup-header"]').contains("Sign up")
    cy.get('input').should('have.length', 4);
  })
  it('navigate to login',()=>{
    cy.get('[data-test="login-btn"]').click()
    cy.get('[data-test="login-header"]').contains("Login")
    cy.get('input').should('have.length', 3)
  })
})