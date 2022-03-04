describe('Travel',()=>{

      it('has 8 cards', () => {
       
        cy.visit("/product") 
        cy.get('.card').its('length').should('be.eq', 8)
       })
    
       it('has a table', () => {
       
        cy.get('.table').should('be.visible')   
       })

       it('view answers works properly', () => {
       cy.contains('View Answers').click()

       
        cy.get('.anstext').should('exist')

        
       })




})