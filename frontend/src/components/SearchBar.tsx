import type { ChangeEvent } from 'react';
import type { SetURLSearchParams } from 'react-router';

interface SearchBarProps {
  onSearchQueryChange: SetURLSearchParams;
}

function SearchBar({ onSearchQueryChange }: SearchBarProps) {
  function handleOnSearchQueryChange(event: ChangeEvent<HTMLInputElement>) {
    onSearchQueryChange((prev) => {
      prev.set('search', event.target.value);
      return prev;
    });
  }
  return (
    <input
      className="max-w-md"
      type="text"
      placeholder="Search products..."
      onChange={handleOnSearchQueryChange}
    />
  );
}

export default SearchBar;
