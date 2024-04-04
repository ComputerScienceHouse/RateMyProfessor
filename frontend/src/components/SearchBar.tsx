import React, { useState, ChangeEvent, FormEvent } from 'react';
import './SearchBar.css'

interface SearchBarProps {
  onSearch: (query: string) => void;
}

const SearchBar: React.FC<SearchBarProps> = ({ onSearch }) => {
  const [query, setQuery] = useState<string>('');

  const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    setQuery(event.target.value);
  };

  const handleSubmit = (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    onSearch(query);
  };

  return (
    
    <form onSubmit={handleSubmit} className='search-bar'>
    <input 
      type="search" 
      name="search" 
      pattern=".*\S.*" 
      required 
      value={query}
      onChange={handleChange}
      />
    <button className="search-btn" type="submit">
      <span>Search</span>
    </button>
  </form>
  );
};

export default SearchBar;
