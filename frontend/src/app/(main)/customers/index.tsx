"use client";

import { Eye, Search } from "lucide-react";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import { GetCustomers } from "./requests";
import { CustomerData, PageRequest } from "./_dto/customers_dto";

const CustomerComponent = () => {

  const router = useRouter();

  const [customers, setcustomer] = useState<CustomerData[]>([]);
  const [search, setSearch] = useState<string>("");
  const [totalPages, setTotalPages] = useState<number>();

  const fetchCustomerList = async (req: PageRequest) => {
    const result = await GetCustomers(req);
    if (result.page_data.total_pages === 0) {
      result.page_data.total_pages = 1;
      setTotalPages(result.page_data.total_pages)
    } else {
      setTotalPages(result.page_data.total_pages)
    }

    setcustomer(result.data);
  };

  const handlePageChange = (page: number) => {
    fetchCustomerList({
      page_number: page,
      page_size: 10,
      search: search,
      sort_by: "",
      sort_direction: "desc",
    });
  };

  useEffect(() => {
    handlePageChange(0)
  }, [search]);

  return (
    <div className="">
      <h2 className="font-bold text-lg text-gray-500">Customer list</h2>
      <p className="text-sm text-gray-500">Customers Data Table. </p>
      <div className="px-3 mt-4">
    
        <form className="mx-auto flex justify-end mb-4">
          <label
            htmlFor={"default-search"}
            className="mb-2 text-sm font-medium text-gray-900 sr-only dark:text-white"
          >
            Search
          </label>
          <div className="relative">
            <div className="absolute inset-y-0 start-0 flex items-center ps-3 pointer-events-none">
              <Search className="w-4 h-4 text-gray-500 dark:text-gray-40" />
            </div>
            <input
              type="search"
              value={search} onChange={(e) => setSearch(e.target.value)}
              id="default-search"
              className="block focus:border-1  border-1 outline-0 border-gray-200 focus:border-[#3a5a40] text-gray-800   w-full p-3 ps-8 text-sm rounded-lg bg-white    dark:placeholder-gray-400 "
              placeholder="Search...."
              required
            />
          </div>
        </form>

        <table className="min-w-full bg-white rounded-lg ">
          <thead className="bg-gradient-to-r from-[#588157] via-[#a3b18a] to-[#588157] rounded-lg text-white">
            <tr>
              <th className="py-3 px-6 text-left rounded-tl-lg ">Name</th>
              <th className="py-3 px-6 text-left">Email</th>
              <th className="py-3 px-6 text-left">Phone Number</th>
              <th className="py-3 px-6 text-left">Account Number</th>
              <th className="py-3 px-6 text-left">Created</th>
              <th className="py-3 px-6 text-left ">Status</th>
              <th className="py-3 px-6 text-left rounded-tr-lg ">Action</th>
            </tr>
          </thead>
          <tbody className="text-gray-700">
            {customers.map((item, index) => (
              <tr className=" hover:bg-indigo-50 transition" key={index}>
                <td className="py-3 px-6">{item.name}</td>
                <td className="py-3 px-6">{item.email}</td>
                <td className="py-3 px-6">{item.phone_number}</td>
                <td className="py-3 px-6">{item.account_number}</td>
                <td className="py-3 px-6">{item.created_at}</td>
                <td className="py-3 px-6">
                  <span className="inline-block px-2 py-1 text-xs font-semibold text-green-700 bg-green-100 rounded-full">
                    Active
                  </span>
                </td>
                <td className="py-3 px-6">
                  <button
                    className="cursor-pointer  px-2 py-1  text-[#a3b18a] bg-[#f5ebe0] rounded-full"
                    onClick={() => router.push(`/customers/${item.id}`)}
                  >
                    <Eye />
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
        <div className="flex items-center space-x-2 mt-2 justify-end">
          <button
            className="px-3 py-1 text-sm text-gray-500 bg-white rounded-md shadow hover:bg-gray-100 disabled:opacity-50"
            disabled
          >
            Prev
          </button>
         
          {new Array(totalPages).fill(0).map((_, index) => (
            <button
              key={index}
              onClick={() => handlePageChange(index)}
              className="cursor-pointer px-3 py-1 text-sm text-gray-600 bg-white hover:bg-[#b5c1a1] hover:text-white rounded-md shadow"
            >
              {index + 1}
            </button>
          ))}

          <button className="px-3 py-1 text-sm text-gray-500 bg-white rounded-md shadow hover:bg-gray-100">
            Next
          </button>
        </div>
      </div>
    </div>
  );
};

export default CustomerComponent;
