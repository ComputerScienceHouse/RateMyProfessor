import { useOidcAccessToken, useOidc, useOidcIdToken } from '@axa-fr/react-oidc'
import React, { useState }  from 'react'
import { Link } from 'react-router-dom'
import Authenticating from '../callbacks/Authenticating'
import AuthenticationError from '../callbacks/AuthenticationError'
import SessionLost from '../callbacks/SessionLost'
import UserInfo from '../UserInfo'
import SearchBar from '../components/SearchBar'


const Home = () => {
    // important hooks
    // const { accessTokenPayload } = useOidcAccessToken()   // this contains the user info in raw json format
    // const userInfo = accessTokenPayload as UserInfo       // 
    // const { idToken, idTokenPayload } = useOidcIdToken()  // this is how you get the users id token
    // const { login, logout, isAuthenticated } = useOidc()  // this gets the functions to login and logout and the logout state

    const [searchResults, setSearchResults] = useState<string[]>([]);

    const handleSearch = (query: string) => {
        // Perform search operation here (e.g., fetch data from API)
        console.log('Searching for:', query);
        // Placeholder for search results
        const searchResults: string[] = []; // Replace this with your actual search results
        // Update searchResults state with search results
        setSearchResults(searchResults);
    };

    return (
        <div>
            <h1 className="display-3">Hello World!</h1>
            <SearchBar onSearch={handleSearch} />
      {/* Display search results */}
      <ul>
        {searchResults.map((result, index) => (
          <li key={index}>{result}</li>
        ))}
      </ul>
        </div>
    )
}

export default Home