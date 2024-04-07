import React from 'react'
import {
    Collapse,
    Container,
    Nav,
    Navbar,
    NavbarToggler,
    NavItem,
} from 'reactstrap'
import { NavLink, useLocation } from 'react-router-dom'
import Profile from './Profile'
import ThemeToggle from './ThemeToggle'

const NavBar: React.FunctionComponent = () => {
    const [isOpen, setIsOpen] = React.useState<boolean>(false)
    const location = useLocation()

    const toggle = () => {
        setIsOpen(!isOpen)
    }

    console.log(location.pathname)

    return (
        <div>
            <Navbar color='primary' dark expand='lg' fixed='top'>
                <Container>
                    <NavLink to='/' className={'navbar-brand'}>
                        Rate My Professors
                    </NavLink>
                    <NavbarToggler onClick={toggle} />
                    <Collapse isOpen={isOpen} navbar>
                        <Nav navbar>
                            {
                                // to add stuff to the navbar, add a NavItem tag with a NavLink to the route
                            }
                            {location.pathname !== "/" &&
                                <NavItem><form className="form-inline my-2 my-lg-0">
                                    <input className="form-control mr-sm-2" type="text" placeholder="Search" />
                                    <button className="btn btn-secondary my-2 my-sm-0" type="submit">Search</button>
                                </form></NavItem>
                            }

                        </Nav>
                        <Nav navbar className='ml-auto'>
                            <Profile />
                        </Nav>
                        <ThemeToggle />
                    </Collapse>
                </Container>
            </Navbar>
        </div>
    )
}

export default NavBar