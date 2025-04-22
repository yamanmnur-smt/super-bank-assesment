"use client";

import { CircleDollarSign, CreditCard, Eye } from "lucide-react";
import Image from "next/image";
import { useEffect, useState } from "react";
import { CustomerDetailData } from "../_dto/customers_dto";
import { GetCustomerDetail } from "../requests";
import { useParams  } from "next/navigation";

const CustomerDetailComponent = () => {
  const {id} = useParams<{id : string}>()

  const [customer, setCustomer] = useState<CustomerDetailData>()

  const getDetail = async ()  =>  {
    const result = await GetCustomerDetail(parseInt(id))
    if(result.data) {

      setCustomer(result.data)
    }
  }

  useEffect(() => {
    getDetail()
  },[])

  return (
    <div className="">
      <h2 className="font-bold text-lg text-gray-500">Customer Detail</h2>
      <p className="text-sm text-gray-500">Customers Detail. </p>
      <div className="h-full p-4 mt-4">
        <div className="flex flex-row space-x-2">
          <div className="bg-[#dad7cd] w-1/2 flex flex-row shadow-md backdrop-blur-sm px-4 py-4 h-20 space-x-2 rounded-xl">
            <div className="px-3 py-2 bg-[#b5c1a1] backdrop-blur-2xl w-15 rounded-xl">
              <CreditCard className="text-[#a3b18a]" size={35} color="white" />
            </div>
            <div className="flex flex-col">
              <span className="text-[#a3b18a] text-sm font-bold">
                Total Balance
              </span>
              <span className="text-[#a3b18a] text-lg font-bold">
                {customer?.total_balance}
              </span>
            </div>
          </div>
          <div className="bg-[#dad7cd] w-1/2 flex flex-row shadow-md backdrop-blur-sm px-4 py-4 h-20 space-x-2 rounded-xl">
            <div className="px-3 py-2 bg-[#b5c1a1] backdrop-blur-2xl w-15 rounded-xl">
              <CreditCard className="text-[#a3b18a]" size={35} color="white" />
            </div>
            <div className="flex flex-col">
              <span className="text-[#a3b18a] text-sm font-bold">
                Total Balance Pockets
              </span>
              <span className="text-[#a3b18a] text-lg font-bold">
                {customer?.total_pockets}
              </span>
            </div>
          </div>
          <div className="bg-[#dad7cd] w-1/2 flex flex-row shadow-md backdrop-blur-sm px-4 py-4 h-20 space-x-2 rounded-xl">
            <div className="px-3 py-2 bg-[#b5c1a1] backdrop-blur-2xl w-15 rounded-xl">
              <CircleDollarSign
                className="text-[#a3b18a]"
                size={35}
                color="white"
              />
            </div>
            <div className="flex flex-col">
              <span className="text-[#a3b18a] text-sm font-bold">
                Total Deposits
              </span>
              <span className="text-[#a3b18a] text-lg font-bold">
                {customer?.total_deposits}
              </span>
            </div>
          </div>
          <div className="bg-[#dad7cd] w-1/2 flex flex-row shadow-md backdrop-blur-sm px-4 py-4 h-20 space-x-2 rounded-xl">
            <div className="px-3 py-2 bg-[#b5c1a1] backdrop-blur-2xl w-15 rounded-xl">
              <CircleDollarSign
                className="text-[#a3b18a]"
                size={35}
                color="white"
              />
            </div>
            <div className="flex flex-col">
              <span className="text-[#a3b18a] text-sm font-bold">
                Register At
              </span>
              <span className="text-[#a3b18a] text-lg font-bold">
                {customer?.created_at}
              </span>
            </div>
          </div>
        </div>
        <div className="flex flex-row space-x-2 mt-3">
          <div className="  bg-white/20 w-1/3 shadow-lg rounded-xl p-6 space-y-2">
            <h2 className="text-lg font-semibold text-gray-500">
              Customer Pockets
            </h2>

            <div className="overflow-y-auto flex-col h-full space-y-2">
              {customer?.pockets.map((item, index) => (

                <div key={index} className="bg-[#dad7cd]  shadow-md backdrop-blur-sm px-3 py-3 h-20 rounded-xl">
                  <div className="flex flex-row space-x-2">

                    <div className="px-3 py-3 bg-[#b5c1a1] backdrop-blur-2xl w-15 rounded-xl">
                      <CircleDollarSign
                        className="text-[#a3b18a]"
                        size={35}
                        color="white"
                      />
                    </div>
                    <div className="flex flex-col mt-2">
                      <span className="text-[#a3b18a] text-sm font-bold">
                        Pocket : {item.name}
                      </span>
                      <span className="text-[#a3b18a] text-lg font-bold">
                        {item.balance}
                      </span>
                    </div>
                  </div>
                </div>
                )
              )}
							
            </div>
          </div>

          <div className="  bg-white/20 w-1/3 shadow-lg rounded-xl p-6 space-y-2">
            <h2 className="text-lg font-semibold text-gray-500">
              Primary Bank Account Info
            </h2>

            {[
              { label: "Name", value: customer?.name },
              { label: "Account Number", value: customer?.banks.account_number },
              { label: "CVC", value: customer?.banks.cvc },
              { label: "Card Number", value: customer?.banks.card_number },
              { label: "Balance", value: customer?.banks.balance },
            ].map((item, idx) => (
              <div
                key={idx}
                className="flex justify-between items-center py-2 border-b  border-b-gray-300"
              >
                <p className="text-sm text-gray-500">{item.label}</p>
                <div className="flex items-center space-x-2">
                  <p className="text-sm font-medium text-gray-800">
                    {item.value}
                  </p>
                </div>
              </div>
            ))}
          </div>

          <div className="  bg-white/20 w-1/3 shadow-lg rounded-xl p-6 space-y-4">
            <h2 className="text-xl font-semibold">Personal Info</h2>

            <div className="flex items-center justify-between border-gray-200 pb-4">
              <div>
                <p className="text-sm text-gray-500">Photo</p>
                <p className="text-xs text-gray-400">
                  150×150px JPEG, PNG Image
                </p>
              </div>
              <div className="relative w-16 h-16">
                <Image
								src="http://localhost:3001/api/v1/file/get-customer-photo?filename=Screenshot%202025-04-17%20102813.png" // replace with your image path
								alt="Profile"
								className="rounded-full border-2 border-green-400"
								layout="fill"
								objectFit="cover"
							/>
            
                <button className="absolute top-0 right-0 bg-white rounded-full p-1 shadow">
                  {/* <XIcon size={12} /> */}
                </button>
              </div>
            </div>

            {[
              { label: "Name", value: customer?.name },
              { label: "Email", value: customer?.email },
              { label: "Phone", value: customer?.phone },
              { label: "Company", value: customer?.company_name },
              { label: "Jobs", value: customer?.jobs },
              { label: "Position", value: customer?.position },
              
              { label: "Birthday", value: "28 May 1996" },
              { label: "Gender", value: customer?.gender },
              { label: "Account Purpose", value: customer?.account_purpose },
              { label: "Address Company", value: customer?.address_company },
            ].map((item, idx) => (
              <div
                key={idx}
                className="flex justify-between items-center py-1 border-b  border-b-gray-500"
              >
                <p className="text-sm text-gray-600">{item.label}</p>
                <div className="flex items-center space-x-2">
                  <p className="text-sm font-medium text-gray-800">
                    {item.value}
                  </p>
                </div>
              </div>
            ))}

            <div className="flex justify-between items-center pt-2">
              <p className="text-sm text-gray-600">Address</p>
              <div className="flex items-center space-x-2">
                <p className="text-sm font-medium text-gray-800">{customer?.address}</p>
              </div>
            </div>
          </div>
        </div>
        <div className="mt-5">
          <h3 className="font-bold text-lg text-gray-500">Deposits Info</h3>

          <table className="min-w-full bg-white rounded-lg mt-3">
            <thead className="bg-gradient-to-r from-[#588157] via-[#a3b18a] to-[#588157] rounded-lg text-white">
              <tr>
                <th className="py-2 px-3 text-left rounded-tl-lg ">Term</th>
                <th className="py-2 px-3 text-left">Amount</th>
                <th className="py-2 px-3 text-left">Interest Rate</th>
                <th className="py-2 px-3 text-left">Start Date</th>
                <th className="py-2 px-3 text-left">Maturity Date</th>
                <th className="py-2 px-3 text-left">Extension Instructions</th>
                <th className="py-2 px-3 text-left ">Status</th>
                <th className="py-2 px-3 text-left rounded-tr-lg ">Action</th>
              </tr>
            </thead>
            <tbody className="text-gray-700">
              {customer?.banks.deposits.map((item, index) => (
                <tr className=" hover:bg-indigo-50 transition" key={index}>
                  <td className="py-3 px-6">{item.term_deposits_types.name}</td>
                  <td className="py-3 px-6">{item.amount}</td>
                  <td className="py-3 px-6">{item.interest_rate}</td>
                  <td className="py-3 px-6">{item.start_date}</td>
                  <td className="py-3 px-6">{item.maturity_date}</td>
                  <td className="py-3 px-6">{item.extension_instructions}</td>
                  <td className="py-3 px-6">
                    <span className="inline-block px-2 py-1 text-xs font-semibold text-green-700 bg-green-100 rounded-full">
                      {item.status}
                    </span>
                  </td>
                  <td className="py-3 px-6">
                    <button className="cursor-pointer  px-2 py-1  text-[#a3b18a] bg-[#f5ebe0] rounded-full">
                      <Eye />
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
};

export default CustomerDetailComponent;
