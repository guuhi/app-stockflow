import type { SetURLSearchParams } from 'react-router';

interface PaginationBarProps {
  currentPage: number;
  totalPages: number;
  onPageChange: SetURLSearchParams;
}

function PaginationBar({
  currentPage,
  totalPages,
  onPageChange,
}: PaginationBarProps) {
  const pages: number[] = [];
  for (let page = 1; page <= totalPages; page++) {
    pages.push(page);
  }

  function handleOnPageChange(page: number) {
    onPageChange((prev) => {
      if (page < 1) {
        page = 1;
      }

      if (page > totalPages) {
        page = totalPages;
      }

      prev.set('page', page.toString());

      return prev;
    });
  }

  return (
    <div className="flex items-center justify-center gap-2">
      <button
        className="px-2 py-1 rounded-md border border-slate-300 disabled:opacity-50"
        disabled={currentPage === 1}
        onClick={() => handleOnPageChange(currentPage - 1)}
      >
        &lt;
      </button>

      {pages.map((page) => (
        <button
          key={page}
          className={`px-3 py-1 rounded-md border border-slate-300 ${
            page === currentPage ? 'bg-indigo-600 text-white' : ''
          }`}
          onClick={() => handleOnPageChange(page)}
        >
          {page}
        </button>
      ))}

      <button
        className="px-2 py-1 rounded-md border border-slate-300 disabled:opacity-50"
        disabled={currentPage === totalPages}
        onClick={() => handleOnPageChange(currentPage + 1)}
      >
        &gt;
      </button>
    </div>
  );
}

export default PaginationBar;
