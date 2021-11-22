<template>

  <div class="h-full w-full flex overflow-hidden antialiased text-gray-800 bg-white">
    <!-- section body side nav -->
    <SideBar />
    <div class="flex-1 flex flex-col">
      <!-- section body top nav -->
      <TopBar />
      <!-- section body header -->
      <Header />
      <!-- main content -->
      <main class="flex-grow flex min-h-0 border-t">
        <!-- section update to tickets -->
        <Ticket />
        <!-- section content -->
        <section aria-label="main content" class="flex min-h-0 flex-col flex-auto border-l">
          <!-- content navigation -->
          <content-nav></content-nav>

          <!-- content caption -->
          <header class="bg-white border-t flex items-center py-1 px-4">
            <div class="flex">
              <h2 id="content-caption" class="font-semibold">
                Tickets requiring your attention (6)
              </h2>
              <span class="ml-3 group relative">
              <button role="details" aria-controls="info-popup" class="text-blue-700 border-b border-dotted border-blue-700 focus:outline-none text-sm">
                What is this?
              </button>
              <div role="tooltip" id="info-popup" class="absolute pt-1 rounded-md rounded-t-lg right-0 transform translate-x-40 mx-auto hidden group-hover:block z-50">
                <div class="border rounded-md rounded-t-lg shadow-lg bg-white w-full max-w-xs w-screen">
                  <header class="font-semibold rounded-t-lg bg-gray-300 px-4 py-2">
                    People are waiting for replies!
                  </header>
                  <div class="p-4 border-t">
                    <p class="mb-4">
                      These are new or open tickets that are assigned to
                      you, unassinged in your group(s) or not assigned to
                      any group.
                    </p>
                    <p class="mb-1">
                      They are ordered by priority and requester update date
                      (oldest first).
                    </p>
                  </div>
                </div>
              </div>
            </span>
            </div>
            <div class="ml-auto">
              <button title="See available tickets in this view" aria-label="play" class="border rounded-md px-3 py-2 leading-none">
                Play
              </button>
            </div>
          </header>

          <!-- content overflow section
                  remove table and thead but keep tbody and change tbody to section, in order
                  to have scrollable overflow section -->
          <table aria-describedby="info-popup" aria-label="open tickets" class="border-t w-full min-h-0 h-full flex flex-col">
            <thead class="flex w-full flex-col px-4">
            <tr class="border-b flex">
              <th class="font-semibold text-left py-3 pl-3 pr-1 w-24">
                <input type="checkbox" name="" id="" />
              </th>
              <th class="font-semibold text-left py-3 px-1 w-24 truncate">
                ID
              </th>
              <th class="font-semibold text-left py-3 px-1 w-full max-w-xs xl:max-w-lg truncate">
                Subject
              </th>
              <th class="font-semibold text-left py-3 px-1 flex-1 truncate">
                Requester
              </th>
              <th class="font-semibold text-left py-3 px-1 flex-1 truncate">
                Requester updated
              </th>
              <th class="font-semibold text-left py-3 px-1 flex-1 truncate">
                Group
              </th>
              <th class="font-semibold text-left py-3 px-1 flex-1 truncate">
                Assignee
              </th>
            </tr>
            </thead>
            <tbody class="flex w-full flex-col flex-1 min-h-0 overflow-hidden px-4">
            <!-- low -->
            <tr class="border-b flex">
              <th class="bg-gray-100 text-left px-3 py-2 flex-1" colspan="7">
                <h2 class="text-sm">
                  <span class="font-normal mr-1">Priority</span><span>Low</span>
                </h2>
              </th>
            </tr>
            <tr role="row" class="hover:bg-blue-100 border-b flex cursor-pointer">
              <td role="cell" headers="select" class="py-3 pl-3 pr-1 w-24 flex items-start">
                <input class="mt-1" type="checkbox" />
                <div class="ml-auto relative group">
                  <span style="
                          padding: 2px 5px;
                          font-size: 0.7rem;
                          position: relative;
                          bottom: 2px;
                        " class="font-mono rounded-sm bg-red-600 text-white leading-none">O</span>
                  <!-- dropdown -->
                  <span class="hidden group-hover:block ml-4 mt-10 w-screen max-w-lg absolute top-0 border shadow-lg p-6 bg-white rounded-md z-50 text-gray-900">
                    <article>
                      <header>
                        <div>
                          <span class="px-3 py-1 uppercase text-xs leading-none rounded-sm bg-red-600 text-white">Open</span>
                          <span class="ml-2 text-gray-700">Incident #12534</span>
                          <span class="ml-1">(Low)</span>
                        </div>
                      </header>
                      <section class="mt-5">
                        <h1 class="text-sm font-semibold mt-3">
                          Quo laudantium error corporis accusamus unde,
                          labore quidem non officiis.
                        </h1>
                        <p class="mt-3">
                          Hi Team,
                          <br />
                          Lorem ipsum dolor sit amet consectetur adipisicing
                          elit. Error accusantium molestias fugit commodi
                          doloremque. <br /><br />
                          Lorem ipsum dolor sit amet consectetur,
                          adipisicing elit? ...
                        </p>
                      </section>
                      <footer class="mt-4">
                        <p class="text-gray-600">Latest Comments</p>
                        <hr class="mt-1" />
                        <div class="flex mt-3">
                          <p class="font-semibold">Nico Braun</p>
                          <p class="ml-auto text-gray-700 text-sm">
                            Yesterday 10:33
                          </p>
                        </div>
                        <p class="mt-2">
                          Dolore odio error inventore sint et dolorum
                          asperiores exercitationem, quisquam esse.
                        </p>
                      </footer>
                    </article>
                  </span>
                  <!-- end dropdown -->
                </div>
              </td>
              <td class="py-3 px-1 w-24">
                #12534
              </td>
              <td class="py-3 px-1 w-full max-w-xs xl:max-w-lg">
                <div class="relative group w-full">
                  <p class="w-full truncate">
                    Quo laudantium error corporis accusamus unde, labore
                    quidem non officiis.
                  </p>
                  <!-- dropdown -->
                  <span class="hidden group-hover:block ml-4 mt-10 w-screen max-w-lg absolute top-0 border shadow-lg p-6 bg-white rounded-md z-50 text-gray-900">
                    <article>
                      <header>
                        <div>
                          <span class="px-3 py-1 uppercase text-xs leading-none rounded-sm bg-red-600 text-white">Open</span>
                          <span class="ml-2 text-gray-700">Incident #12534</span>
                          <span class="ml-1">(Low)</span>
                        </div>
                      </header>
                      <section class="mt-5">
                        <h1 class="text-sm font-semibold mt-3">
                          Quo laudantium error corporis accusamus unde,
                          labore quidem non officiis.
                        </h1>
                        <p class="mt-3">
                          Hi Team,
                          <br />
                          Lorem ipsum dolor sit amet consectetur adipisicing
                          elit. Error accusantium molestias fugit commodi
                          doloremque. <br /><br />
                          Lorem ipsum dolor sit amet consectetur,
                          adipisicing elit? ...
                        </p>
                      </section>
                      <footer class="mt-4">
                        <p class="text-gray-600">Latest Comments</p>
                        <hr class="mt-1" />
                        <div class="flex mt-3">
                          <p class="font-semibold">Nico Braun</p>
                          <p class="ml-auto text-gray-700 text-sm">
                            Yesterday 10:33
                          </p>
                        </div>
                        <p class="mt-2">
                          Dolore odio error inventore sint et dolorum
                          asperiores exercitationem, quisquam esse.
                        </p>
                      </footer>
                    </article>
                  </span>
                  <!-- end dropdown -->
                </div>
              </td>
              <td class="py-3 px-1 flex-1 truncate">
                Marla Darsuz
              </td>
              <td class="py-3 px-1 flex-1 truncate">
                Tuesday 09:56
              </td>
              <td class="py-3 px-1 flex-1 truncate">
                UK Support
              </td>
              <td class="py-3 px-1 flex-1 truncate">
                Nico Braun
              </td>
            </tr>
            <tr role="row" class="hover:bg-blue-100 border-b flex cursor-pointer">
              <td role="cell" headers="select" class="py-3 pl-3 pr-1 w-24 flex items-start">
                <input class="mt-1" type="checkbox" />
                <div class="ml-auto relative group">
                  <span style="
                          padding: 2px 5px;
                          font-size: 0.7rem;
                          position: relative;
                          bottom: 2px;
                        " class="font-mono rounded-sm bg-red-600 text-white leading-none">O</span>
                  <!-- dropdown -->
                  <span class="hidden group-hover:block ml-4 mt-10 w-screen max-w-lg absolute top-0 border shadow-lg p-6 bg-white rounded-md z-50 text-gray-900">
                    <article>
                      <header>
                        <div>
                          <span class="px-3 py-1 uppercase text-xs leading-none rounded-sm bg-red-600 text-white">Open</span>
                          <span class="ml-2 text-gray-700">Incident #12534</span>
                          <span class="ml-1">(Low)</span>
                        </div>
                      </header>
                      <section class="mt-5">
                        <h1 class="text-sm font-semibold mt-3">
                          Lorem ipsum dolor sit amet, consectetur
                          adipisicing elit.
                        </h1>
                        <p class="mt-3">
                          Hi, <br />
                          <br />
                          Lorem ipsum dolor sit amet consectetur adipisicing
                          elit. Error accusantium molestias fugit commodi
                          doloremque. <br /><br />
                          Lorem ipsum dolor sit amet consectetur,
                          adipisicing elit. Odit consequatur natus aut
                          reiciendis nisi sed! Temporibus, quibusdam
                          voluptates? ...
                        </p>
                      </section>
                      <footer class="mt-4">
                        <p class="text-gray-600">Latest Comments</p>
                        <hr class="mt-1" />
                        <div class="flex mt-3">
                          <p class="font-semibold">Nico Braun</p>
                          <p class="ml-auto text-gray-700 text-sm">
                            Today 13:30
                          </p>
                        </div>
                        <p class="mt-2">
                          Dolore odio error inventore sint et dolorum
                          asperiores exercitationem, quisquam esse tempora
                          aliquam voluptates quibusdam quae debitis
                          laboriosam iure ea quos.
                        </p>
                      </footer>
                    </article>
                  </span>
                  <!-- end dropdown -->
                </div>
              </td>
              <td class="py-3 px-1 w-24">
                #12534
              </td>
              <td class="py-3 px-1 w-full max-w-xs xl:max-w-lg">
                <div class="relative group w-full">
                  <p class="w-full truncate">
                    Lorem ipsum dolor sit amet, consectetur adipisicing
                    elit.
                  </p>
                  <!-- dropdown -->
                  <span class="hidden group-hover:block ml-4 mt-10 w-screen max-w-lg absolute top-0 border shadow-lg p-6 bg-white rounded-md z-50 text-gray-900">
                    <article>
                      <header>
                        <div>
                          <span class="px-3 py-1 uppercase text-xs leading-none rounded-sm bg-red-600 text-white">Open</span>
                          <span class="ml-2 text-gray-700">Incident #12534</span>
                          <span class="ml-1">(Low)</span>
                        </div>
                      </header>
                      <section class="mt-5">
                        <h1 class="text-sm font-semibold mt-3">
                          Lorem ipsum dolor sit amet, consectetur
                          adipisicing elit.
                        </h1>
                        <p class="mt-3">
                          Hi, <br />
                          <br />
                          Lorem ipsum dolor sit amet consectetur adipisicing
                          elit. Error accusantium molestias fugit commodi
                          doloremque. <br /><br />
                          Lorem ipsum dolor sit amet consectetur,
                          adipisicing elit. Odit consequatur natus aut
                          reiciendis nisi sed! Temporibus, quibusdam
                          voluptates? ...
                        </p>
                      </section>
                      <footer class="mt-4">
                        <p class="text-gray-600">Latest Comments</p>
                        <hr class="mt-1" />
                        <div class="flex mt-3">
                          <p class="font-semibold">Nico Braun</p>
                          <p class="ml-auto text-gray-700 text-sm">
                            Today 13:30
                          </p>
                        </div>
                        <p class="mt-2">
                          Dolore odio error inventore sint et dolorum
                          asperiores exercitationem, quisquam esse tempora
                          aliquam voluptates quibusdam quae debitis
                          laboriosam iure ea quos.
                        </p>
                      </footer>
                    </article>
                  </span>
                  <!-- end dropdown -->
                </div>
              </td>
              <td class="py-3 px-1 flex-1 truncate">
                Jim Beam
              </td>
              <td class="py-3 px-1 flex-1 truncate">
                Friday 10:29
              </td>
              <td class="py-3 px-1 flex-1 truncate">
                UK Support
              </td>
              <td class="py-3 px-1 flex-1 truncate">
                Nico Braun
              </td>
            </tr>
            <tr role="row" class="hover:bg-blue-100 border-b flex cursor-pointer">
              <td role="cell" headers="select" class="py-3 pl-3 pr-1 w-24 flex items-start">
                <input class="mt-1" type="checkbox" />
                <div class="ml-auto relative group">
                  <span style="
                          padding: 2px 5px;
                          font-size: 0.7rem;
                          position: relative;
                          bottom: 2px;
                        " class="font-mono rounded-sm bg-yellow-400 text-black leading-none">N</span>
                  <!-- dropdown -->
                  <span class="hidden group-hover:block ml-4 mt-10 w-screen max-w-lg absolute top-0 border shadow-lg p-6 bg-white rounded-md z-50 text-gray-900">
                    <article>
                      <header>
                        <div>
                          <span class="px-3 py-1 uppercase text-xs leading-none rounded-sm bg-yellow-400 text-black">New</span>
                          <span class="ml-2 text-gray-700">Incident #12534</span>
                          <span class="ml-1">(Low)</span>
                        </div>
                      </header>
                      <section class="mt-5">
                        <h1 class="text-sm font-semibold mt-3">
                          Excepturi at labore vel accusamus exercitationem
                          quam, amet provident!
                        </h1>
                        <p class="mt-3">
                          Lorem ipsum dolor sit amet consectetur,
                          adipisicing elit. Odit consequatur natus aut
                          reiciendis nisi sed! Temporibus, quibusdam
                          voluptates? ...
                        </p>
                      </section>
                      <footer class="mt-4 hidden">
                        <p class="text-gray-600">Latest Comments</p>
                        <hr class="mt-1" />
                        <div class="flex mt-3">
                          <p class="font-semibold"></p>
                          <p class="ml-auto text-gray-700 text-sm"></p>
                        </div>
                        <p class="mt-2"></p>
                      </footer>
                    </article>
                  </span>
                  <!-- end dropdown -->
                </div>
              </td>
              <td class="py-3 px-1 w-24">
                #12534
              </td>
              <td class="py-3 px-1 w-full max-w-xs xl:max-w-lg">
                <div class="relative group w-full">
                  <p class="w-full truncate">
                    Excepturi at labore vel accusamus exercitationem quam,
                    amet provident!
                  </p>
                  <!-- dropdown -->
                  <span class="hidden group-hover:block ml-4 mt-10 w-screen max-w-lg absolute top-0 border shadow-lg p-6 bg-white rounded-md z-50 text-gray-900">
                    <article>
                      <header>
                        <div>
                          <span class="px-3 py-1 uppercase text-xs leading-none rounded-sm bg-yellow-400 text-black">New</span>
                          <span class="ml-2 text-gray-700">Incident #12534</span>
                          <span class="ml-1">(Low)</span>
                        </div>
                      </header>
                      <section class="mt-5">
                        <h1 class="text-sm font-semibold mt-3">
                          Excepturi at labore vel accusamus exercitationem
                          quam, amet provident!
                        </h1>
                        <p class="mt-3">
                          Lorem ipsum dolor sit amet consectetur,
                          adipisicing elit. Odit consequatur natus aut
                          reiciendis nisi sed! Temporibus, quibusdam
                          voluptates? ...
                        </p>
                      </section>
                      <footer class="mt-4 hidden">
                        <p class="text-gray-600">Latest Comments</p>
                        <hr class="mt-1" />
                        <div class="flex mt-3">
                          <p class="font-semibold"></p>
                          <p class="ml-auto text-gray-700 text-sm"></p>
                        </div>
                        <p class="mt-2"></p>
                      </footer>
                    </article>
                  </span>
                  <!-- end dropdown -->
                </div>
              </td>
              <td class="py-3 px-1 flex-1 truncate">
                Paul Ferrez
              </td>
              <td class="py-3 px-1 flex-1 truncate">
                Today 13:45
              </td>
              <td class="py-3 px-1 flex-1 truncate">
                -
              </td>
              <td class="py-3 px-1 flex-1 truncate">
                -
              </td>
            </tr>
            <tr role="row" class="hover:bg-blue-100 border-b flex cursor-pointer">
              <td role="cell" headers="select" class="py-3 pl-3 pr-1 w-24 flex items-start">
                <input class="mt-1" type="checkbox" />
                <div class="ml-auto relative group">
                  <span style="
                          padding: 2px 5px;
                          font-size: 0.7rem;
                          position: relative;
                          bottom: 2px;
                        " class="font-mono rounded-sm bg-red-600 text-white leading-none">O</span>
                  <!-- dropdown -->
                  <span class="hidden group-hover:block ml-4 mt-10 w-screen max-w-lg absolute top-0 border shadow-lg p-6 bg-white rounded-md z-50 text-gray-900">
                    <article>
                      <header>
                        <div>
                          <span class="px-3 py-1 uppercase text-xs leading-none rounded-sm bg-red-600 text-white">Open</span>
                          <span class="ml-2 text-gray-700">Incident #12534</span>
                          <span class="ml-1">(Low)</span>
                        </div>
                      </header>
                      <section class="mt-5">
                        <h1 class="text-sm font-semibold mt-3">
                          impedit possimus praesentium voluptatum omnis
                          assumenda rem autem magni consequatur nostrum
                          distinctio unde.
                        </h1>
                        <p class="mt-3">
                          Hi Team,
                          <br />
                          Lorem ipsum dolor sit amet consectetur adipisicing
                          elit. Error accusantium molestias fugit commodi
                          doloremque. <br /><br />
                          Lorem ipsum dolor sit amet consectetur,
                          adipisicing elit? ...
                        </p>
                      </section>
                      <footer class="mt-4">
                        <p class="text-gray-600">Latest Comments</p>
                        <hr class="mt-1" />
                        <div class="flex mt-3">
                          <p class="font-semibold">Nico Braun</p>
                          <p class="ml-auto text-gray-700 text-sm">
                            Yesterday 10:33
                          </p>
                        </div>
                        <p class="mt-2">
                          Dolore odio error inventore sint et dolorum
                          asperiores exercitationem, quisquam esse.
                        </p>
                      </footer>
                    </article>
                  </span>
                  <!-- end dropdown -->
                </div>
              </td>
              <td class="py-3 px-1 w-24">
                #12534
              </td>
              <td class="py-3 px-1 w-full max-w-xs xl:max-w-lg">
                <div class="relative group w-full">
                  <p class="w-full truncate">
                    impedit possimus praesentium voluptatum omnis assumenda
                    rem autem magni consequatur nostrum distinctio unde.
                  </p>
                  <!-- dropdown -->
                  <span class="hidden group-hover:block ml-4 mt-10 w-screen max-w-lg absolute top-0 border shadow-lg p-6 bg-white rounded-md z-50 text-gray-900">
                    <article>
                      <header>
                        <div>
                          <span class="px-3 py-1 uppercase text-xs leading-none rounded-sm bg-red-600 text-white">Open</span>
                          <span class="ml-2 text-gray-700">Incident #12534</span>
                          <span class="ml-1">(Low)</span>
                        </div>
                      </header>
                      <section class="mt-5">
                        <h1 class="text-sm font-semibold mt-3">
                          impedit possimus praesentium voluptatum omnis
                          assumenda rem autem magni consequatur nostrum
                          distinctio unde.
                        </h1>
                        <p class="mt-3">
                          Hi Team,
                          <br />
                          Lorem ipsum dolor sit amet consectetur adipisicing
                          elit. Error accusantium molestias fugit commodi
                          doloremque. <br /><br />
                          Lorem ipsum dolor sit amet consectetur,
                          adipisicing elit? ...
                        </p>
                      </section>
                      <footer class="mt-4">
                        <p class="text-gray-600">Latest Comments</p>
                        <hr class="mt-1" />
                        <div class="flex mt-3">
                          <p class="font-semibold">Nico Braun</p>
                          <p class="ml-auto text-gray-700 text-sm">
                            Yesterday 10:33
                          </p>
                        </div>
                        <p class="mt-2">
                          Dolore odio error inventore sint et dolorum
                          asperiores exercitationem, quisquam esse.
                        </p>
                      </footer>
                    </article>
                  </span>
                  <!-- end dropdown -->
                </div>
              </td>
              <td class="py-3 px-1 flex-1 truncate">
                Sara Dechicco
              </td>
              <td class="py-3 px-1 flex-1 truncate">
                Thursday 09:22
              </td>
              <td class="py-3 px-1 flex-1 truncate">
                UK Support
              </td>
              <td class="py-3 px-1 flex-1 truncate">
                Nico Braun
              </td>
            </tr>
            <!-- medium -->
            <tr class="border-b flex">
              <th class="bg-gray-100 text-left px-3 py-2 flex-1" colspan="7">
                <h2 class="text-sm">
                  <span class="font-normal mr-1">Priority</span>
                  <span>Medium</span>
                </h2>
              </th>
            </tr>
            <tr role="row" class="hover:bg-blue-100 border-b flex cursor-pointer">
              <td role="cell" headers="select" class="py-3 pl-3 pr-1 w-24 flex items-start">
                <input class="mt-1" type="checkbox" />
                <div class="ml-auto relative group">
                  <span style="
                          padding: 2px 5px;
                          font-size: 0.7rem;
                          position: relative;
                          bottom: 2px;
                        " class="font-mono rounded-sm bg-yellow-400 text-black leading-none">N</span>
                  <!-- dropdown -->
                  <span class="hidden group-hover:block ml-4 mt-10 w-screen max-w-lg absolute top-0 border shadow-lg p-6 bg-white rounded-md z-50 text-gray-900">
                    <article>
                      <header>
                        <div>
                          <span class="px-3 py-1 uppercase text-xs leading-none rounded-sm bg-yellow-400 text-black">New</span>
                          <span class="ml-2 text-gray-700">Incident #12534</span>
                          <span class="ml-1">(Medium)</span>
                        </div>
                      </header>
                      <section class="mt-5">
                        <h1 class="text-sm font-semibold mt-3">
                          Excepturi at labore vel accusamus exercitationem
                          assumenda ex incidunt eum quam, amet provident!
                        </h1>
                        <p class="mt-3">
                          Hi, <br />
                          <br />
                          Lorem ipsum dolor sit amet consectetur adipisicing
                          elit. Error accusantium molestias fugit commodi
                          doloremque. <br /><br />
                          Lorem ipsum dolor sit amet consectetur,
                          adipisicing elit. Odit consequatur natus aut
                          reiciendis nisi sed! Temporibus, quibusdam
                          voluptates? ...
                        </p>
                      </section>
                      <footer class="mt-4 hidden">
                        <p class="text-gray-600">Latest Comments</p>
                        <hr class="mt-1" />
                        <div class="flex mt-3">
                          <p class="font-semibold"></p>
                          <p class="ml-auto text-gray-700 text-sm"></p>
                        </div>
                        <p class="mt-2"></p>
                      </footer>
                    </article>
                  </span>
                  <!-- end dropdown -->
                </div>
              </td>
              <td class="py-3 px-1 w-24">
                #12534
              </td>
              <td class="py-3 px-1 w-full max-w-xs xl:max-w-lg">
                <div class="relative group w-full">
                  <p class="w-full truncate">
                    Excepturi at labore vel accusamus exercitationem
                    assumenda ex incidunt eum quam, amet provident!
                  </p>
                  <!-- dropdown -->
                  <span class="hidden group-hover:block ml-4 mt-10 w-screen max-w-lg absolute top-0 border shadow-lg p-6 bg-white rounded-md z-50 text-gray-900">
                    <article>
                      <header>
                        <div>
                          <span class="px-3 py-1 uppercase text-xs leading-none rounded-sm bg-yellow-400 text-black">New</span>
                          <span class="ml-2 text-gray-700">Incident #12534</span>
                          <span class="ml-1">(Medium)</span>
                        </div>
                      </header>
                      <section class="mt-5">
                        <h1 class="text-sm font-semibold mt-3">
                          Excepturi at labore vel accusamus exercitationem
                          assumenda ex incidunt eum quam, amet provident!
                        </h1>
                        <p class="mt-3">
                          Hi, <br />
                          <br />
                          Lorem ipsum dolor sit amet consectetur adipisicing
                          elit. Error accusantium molestias fugit commodi
                          doloremque. <br /><br />
                          Lorem ipsum dolor sit amet consectetur,
                          adipisicing elit. Odit consequatur natus aut
                          reiciendis nisi sed! Temporibus, quibusdam
                          voluptates? ...
                        </p>
                      </section>
                      <footer class="mt-4 hidden">
                        <p class="text-gray-600"></p>
                        <hr class="mt-1" />
                        <div class="flex mt-3">
                          <p class="font-semibold"></p>
                          <p class="ml-auto text-gray-700 text-sm"></p>
                        </div>
                        <p class="mt-2">
                          Dolore odio error inventore sint et dolorum
                          asperiores exercitationem, quisquam esse tempora
                          aliquam voluptates quibusdam quae debitis
                          laboriosam iure ea quos.
                        </p>
                      </footer>
                    </article>
                  </span>
                  <!-- end dropdown -->
                </div>
              </td>
              <td class="py-3 px-1 flex-1 truncate">
                Freddy Murro
              </td>
              <td class="py-3 px-1 flex-1 truncate">
                Today 12:13
              </td>
              <td class="py-3 px-1 flex-1 truncate">
                -
              </td>
              <td class="py-3 px-1 flex-1 truncate">
                -
              </td>
            </tr>
            <tr role="row" class="hover:bg-blue-100 border-b flex cursor-pointer">
              <td role="cell" headers="select" class="py-3 pl-3 pr-1 w-24 flex items-start">
                <input class="mt-1" type="checkbox" />
                <div class="ml-auto relative group">
                  <span style="
                          padding: 2px 5px;
                          font-size: 0.7rem;
                          position: relative;
                          bottom: 2px;
                        " class="font-mono rounded-sm bg-red-600 text-white leading-none">O</span>
                  <!-- dropdown -->
                  <span class="hidden group-hover:block ml-4 mt-10 w-screen max-w-lg absolute top-0 border shadow-lg p-6 bg-white rounded-md z-50 text-gray-900">
                    <article>
                      <header>
                        <div>
                          <span class="px-3 py-1 uppercase text-xs leading-none rounded-sm bg-red-600 text-white">Open</span>
                          <span class="ml-2 text-gray-700">Incident #12534</span>
                          <span class="ml-1">(Medium)</span>
                        </div>
                      </header>
                      <section class="mt-5">
                        <h1 class="text-sm font-semibold mt-3">
                          Odit consequatur natus aut reiciendis nisi sed!
                        </h1>
                        <p class="mt-3">
                          Hi, <br />
                          <br />
                          Lorem ipsum dolor sit amet consectetur adipisicing
                          elit. Error accusantium molestias fugit commodi
                          doloremque. <br /><br />
                          Lorem ipsum dolor sit amet consectetur,
                          adipisicing elit. Temporibus, quibusdam
                          voluptates? ...
                        </p>
                      </section>
                      <footer class="mt-4">
                        <p class="text-gray-600">Latest Comments</p>
                        <hr class="mt-1" />
                        <div class="flex mt-3">
                          <p aria-label="commenter" class="font-semibold">
                            Nico Braun
                          </p>
                          <p aria-label="time" class="ml-auto text-gray-700 text-sm">
                            Today 06:34
                          </p>
                        </div>
                        <p aria-label="comment" class="mt-2">
                          Quos explicabo sed nisi totam dolores deleniti
                          blanditiis culpa dolor provident perferendis
                          sapiente corrupti repudiandae ea dolore.
                        </p>
                      </footer>
                    </article>
                  </span>
                  <!-- end dropdown -->
                </div>
              </td>
              <td class="py-3 px-1 w-24">
                #12534
              </td>
              <td class="py-3 px-1 w-full max-w-xs xl:max-w-lg">
                <div class="relative group w-full">
                  <p class="w-full truncate">
                    Odit consequatur natus aut reiciendis nisi sed!
                  </p>
                  <!-- dropdown -->
                  <span class="hidden group-hover:block ml-4 mt-10 w-screen max-w-lg absolute top-0 border shadow-lg p-6 bg-white rounded-md z-50 text-gray-900">
                    <article>
                      <header>
                        <div>
                          <span class="px-3 py-1 uppercase text-xs leading-none rounded-sm bg-red-600 text-white">Open</span>
                          <span class="ml-2 text-gray-700">Incident #12534</span>
                          <span class="ml-1">(Medium)</span>
                        </div>
                      </header>
                      <section class="mt-5">
                        <h1 class="text-sm font-semibold mt-3">
                          Odit consequatur natus aut reiciendis nisi sed!
                        </h1>
                        <p class="mt-3">
                          Hi, <br />
                          <br />
                          Lorem ipsum dolor sit amet consectetur adipisicing
                          elit. Error accusantium molestias fugit commodi
                          doloremque. <br /><br />
                          Lorem ipsum dolor sit amet consectetur,
                          adipisicing elit. Temporibus, quibusdam
                          voluptates? ...
                        </p>
                      </section>
                      <footer class="mt-4">
                        <p class="text-gray-600">Latest Comments</p>
                        <hr class="mt-1" />
                        <div class="flex mt-3">
                          <p aria-label="commenter" class="font-semibold">
                            Nico Braun
                          </p>
                          <p aria-label="time" class="ml-auto text-gray-700 text-sm">
                            Today 06:34
                          </p>
                        </div>
                        <p aria-label="comment" class="mt-2">
                          Quos explicabo sed nisi totam dolores deleniti
                          blanditiis culpa dolor provident perferendis
                          sapiente corrupti repudiandae ea dolore.
                        </p>
                      </footer>
                    </article>
                  </span>
                  <!-- end dropdown -->
                </div>
              </td>
              <td class="py-3 px-1 flex-1 truncate">
                Carla Sandaers
              </td>
              <td class="py-3 px-1 flex-1 truncate">
                Today 08:13
              </td>
              <td class="py-3 px-1 flex-1 truncate">
                -
              </td>
              <td class="py-3 px-1 flex-1 truncate">
                -
              </td>
            </tr>
            <!-- high -->
            <tr class="border-b flex hidden">
              <th class="bg-gray-100 text-left px-3 py-2 flex-1" colspan="7">
                <h2 class="text-sm">
                  <span class="font-normal mr-1">Priority</span>
                  <span>High</span>
                </h2>
              </th>
            </tr>
            </tbody>
          </table>

          <!-- content footer, currently hidden -->
          <footer aria-label="content footer" class="flex p-3 bg-white border-t hidden">
            footer
          </footer>
        </section>
      </main>
    </div>
  </div>
</template>

<script>
import TopBar from './topBar.vue'
import SideBar from './sideBar.vue'
import Header from './header.vue'
import Ticket from './ticket.vue'
import ContentNav from './contentNav.vue'
export default {
  components: { ContentNav, Header, SideBar, TopBar, Ticket }
}
</script>

<style lang="scss" scoped>
html {
  font-size: 14px;
  line-height: 1.285;
  font-family: system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI",
  Roboto, Oxygen-Sans, Ubuntu, Cantarell, "Helvetica Neue", Arial,
  sans-serif;
}
html, body {
  width: 100%;
  height: 100%;
}

/* can be configured in tailwind.config.js */
.group:hover .group-hover\:block {
  display: block;
}

.focus\:cursor-text:focus {
  cursor: text;
}

.focus\:w-64:focus {
  width: 16rem;
}

/* zendesk styles */
.h-16 {
  height: 50px;
}
.bg-teal-900 {
  background: #03363d;
}
.bg-gray-100 {
  background: #f8f9f9;
}
.hover\:border-green-500:hover {
  border-color: #78a300;
}
</style>
