"use client";

import { Users, BookLock, Scale, CreditCard, BookKey } from "lucide-react";

interface ChartInterface {
  name : string
  value : number
}

interface BarInterface {
  letter : string
  frequency : number
}

import * as d3 from "d3";
import { useEffect, useRef, useState } from "react";
import { GetDashboard } from "./requests";
import { DashboardTotalCard } from "./_dto/dashboard_dto";

const DashboardComponent = () => {

  const [totalCard, setTotalCard] = useState<DashboardTotalCard>()
  const [pieData, setPieData] = useState<ChartInterface[]>()
  const [barData, setBarData] = useState<BarInterface[]>()

  const barchartref = useRef(null);

  const piechartref = useRef(null);

  const renderPieChart = () => {
    if (pieData == undefined) {
      return ;
    }

    const width = 180;
    const height = Math.min(width, 180);

    const color = d3
      .scaleOrdinal()
      .domain(pieData.map((d) => d.name))
      .range(["#e4e3de", "#a3b18a", "#588157", "#3a5a40", "#344e41"]);

    const pie = d3
      .pie<ChartInterface>()
      .sort(null)
      .value((d) => d.value);
    const radius = Math.min(width, height) / 2 - 1
    const arc = d3
      .arc<d3.PieArcDatum<ChartInterface>>()
      .innerRadius(0)
      .outerRadius(radius);
    
    const labelRadius = radius * 0.8;

    const arcLabel = d3.arc<d3.PieArcDatum<ChartInterface>>().innerRadius(labelRadius).outerRadius(labelRadius);

    const arcs = pie(pieData);

    const svg = d3
      .select(piechartref.current)
      .attr("width", width)
      .attr("height", height)
      .attr("viewBox", [-width / 2, -height / 2, width, height])
      .attr("style", "max-width: 100%; height: auto; font: 10px sans-serif;");

    svg
      .append("g")
      .selectAll()
      .data(arcs)
      .join("path")
      .attr("fill", (d) => color(d.data.name) as string)
      .attr("d", arc)
      .append("title")
      .text((d) => `${d.data.name}: ${d.data.value.toLocaleString("en-US")}`);

   
    svg
      .append("g")
      .attr("text-anchor", "middle")
      .selectAll()
      .data(arcs)
      .join("text")
      .attr("transform", (d) => `translate(${arcLabel.centroid(d)})`)
      .call((text) =>
        text
          .append("tspan")
          .attr("y", "-0.4em")
          .attr("font-weight", "bold")
          .text((d) => d.data.name)
      )
      .call((text) =>
        text
          .filter((d) => d.endAngle - d.startAngle > 0.25)
          .append("tspan")
          .attr("x", 0)
          .attr("y", "0.7em")
          .attr("fill-opacity", 0.7)
          .text((d) => d.data.value.toLocaleString("en-US"))
      );
  };

  const renderChart = () => {
    if (barData === undefined) {
      return ;
    }
    const width = 600;
    const height = 300;
    const marginTop = 30;
    const marginRight = 0;
    const marginBottom = 30;
    const marginLeft = 60;

    d3.select(barchartref.current).selectAll("*").remove();
    const x = d3
      .scaleBand()
      .domain(barData.map((d) => d.letter)) // descending frequency
      .range([marginLeft, width - marginRight])
      .padding(0.2);

    const y = d3
      .scaleLinear()
      .domain([0, d3.max(barData, (d) => d.frequency) ?? 0])
      .range([height - marginBottom, marginTop]);

    const svg = d3
      .select(barchartref.current)
      .attr("width", width)
      .attr("height", height)
      .attr("viewBox", [0, 0, width, height])
      .attr("style", "max-width: 100%; height: auto;");

    const tooltip = d3
      .select("body")
      .append("div")
      .style("position", "absolute")
      .style("background", "rgba(0, 0, 0, 0.7)")
      .style("color", "white")
      .style("padding", "5px 10px")
      .style("border-radius", "5px")
      .style("pointer-events", "none")
      .style("opacity", 0);

    svg
      .append("g")
      .attr("fill", "#a3b18a")
      .selectAll()
      .data(barData)
      .join("rect")
      .attr("x", (d) => x(d.letter) ?? 0)
      .attr("y", (d) => y(d.frequency))
      .attr("height", (d) => y(0) - y(d.frequency))
      .attr("width", x.bandwidth())
      .on("mouseover", (event, d) => {
        tooltip
          .style("opacity", 1)
          .html(`Month: ${d.letter}<br>Frequency: ${d.frequency}`);
      })
      .on("mousemove", (event) => {
        tooltip
          .style("left", `${event.pageX + 10}px`)
          .style("top", `${event.pageY - 20}px`);
      })
      .on("mouseout", () => {
        tooltip.style("opacity", 0);
      });

    svg
      .append("g")
      .attr("transform", `translate(0,${height - marginBottom})`)
      .call(d3.axisBottom(x).tickSizeOuter(0));

    svg
      .append("g")
      .attr("transform", `translate(${marginLeft},0)`)
      .call(
        d3.axisLeft(y).tickFormat((y) => {
          const number = Number(y) * 100;
          const formats = new Intl.NumberFormat("id-ID", {
            maximumSignificantDigits: 3,
          }).format(number);
          return "Rp" + formats;
        })
      )
      .call((g) => g.select(".domain").remove())
      .call((g) =>
        g
          .append("text")
          .attr("x", -marginLeft)
          .attr("y", 10)
          .attr("fill", "currentColor")
          .attr("text-anchor", "start")
          .text("↑ Frequency (%)")
      );
  };

  const fetchDashboard = async () => {
    const res = await GetDashboard()
    const pie_datas = res.data.pie_data.map(item => {
      const data : ChartInterface = {
        name : item.label,
        value : parseInt(item.value.replace(/[Rp.]/g, ''))
      }
      return data
    })
    setPieData(pie_datas)

    const bar_datas = res.data.bar_data.map(item => {
      const data : BarInterface = {
        letter : item.label,
        frequency : parseInt(item.value.replace(/[Rp.]/g, ''))
      }
      return data
    })

    setBarData(bar_datas)

    setTotalCard(res.data.total_card)
  }

  useEffect(() => {
  
    fetchDashboard()
  }, []);

  useEffect(() => {
    renderChart();
    renderPieChart();
  }, [pieData, barData, totalCard]);

  return (
    <div className="">
      <h2 className="font-bold text-lg text-gray-500">Dashboard</h2>
      <p className="text-sm text-gray-500">Dashboard CMS. </p>
      <div className="h-full p-4 mt-4">
        <div className="flex flex-row space-x-2 ">
          <div className="flex flex-col w-full space-y-4">
            <div className="flex flex-row space-x-3">
              <div className="bg-[#dad7cd] w-full flex flex-row shadow-md backdrop-blur-sm px-4 py-4 space-x-2 rounded-xl">
                <div className="px-3 py-3 bg-[#b5c1a1] backdrop-blur-2xl w-15 rounded-xl">
                  <Users className="text-[#a3b18a]" size={35} color="white" />
                </div>
                <div className="flex flex-col mt-2">
                  <span className="text-[#a3b18a] text-sm font-bold">
                    Customers
                  </span>
                  <span className="text-[#a3b18a] text-4xl font-bold">
                    {totalCard?.total_customers}
                  </span>
                </div>
              </div>
              <div className="bg-[#dad7cd] w-full flex flex-row shadow-md backdrop-blur-sm px-4 py-4 space-x-2 rounded-xl">
                <div className="px-3 py-3 bg-[#b5c1a1] backdrop-blur-2xl w-15 rounded-xl">
                  <BookLock className="text-white" size={35} color="white" />
                </div>
                <div className="flex flex-col mt-2">
                  <span className="text-[#a3b18a] text-sm font-bold">
                    Total Deposits
                  </span>
                  <span className="text-[#a3b18a] text-4xl font-bold">
                    {totalCard?.total_deposits}
                  </span>
                </div>
              </div>

              <div className="bg-[#dad7cd] w-full flex flex-row shadow-md backdrop-blur-sm px-4 py-4 space-x-2 rounded-xl">
                <div className="px-3 py-3 bg-[#b5c1a1] backdrop-blur-2xl w-15 rounded-xl">
                  <CreditCard className="text-white" size={35} color="white" />
                </div>
                <div className="flex flex-col mt-2">
                  <span className="text-[#a3b18a] text-sm font-bold">
                    Total Balance
                  </span>
                  <span className="text-[#a3b18a] text-4xl font-bold">
                    {totalCard?.total_balance}
                  </span>
                </div>
              </div>
            </div>
            <div className="flex flex-row">
              <div>

                <svg ref={barchartref}></svg>
              </div>
              <div className="flex justify-center w-1/2">
                <div className="bg-[#dad7cd] shadow-md backdrop-blur-sm p-4 rounded-xl w-full space-x-2 flex flex-row">
                  <svg ref={piechartref} className="items-center"></svg>
                  <div className="information flex flex-col">

                    {pieData?.map((item, idx) => (
                      <div
                        key={idx}
                        className="flex flex-row justify-between items-center py-2 border-b  border-b-gray-300 w-40"
                      >
                        <p className="text-sm text-gray-500">{item.name}</p>
                        <div className="flex items-center space-x-2">
                          <p className="text-sm font-medium text-gray-800">
                            {item.value}
                          </p>
                        </div>
                      </div>
                    ))}
                  </div>
                </div>
              </div>
            </div>
            
          </div>
       
        </div>
      </div>
    </div>
  );
};

export default DashboardComponent;
