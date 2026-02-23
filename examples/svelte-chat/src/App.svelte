<script>
  import { fetchWithPDP } from '../../../libs/js/index.js';

  let messages = $state([
    { role: 'ai', content: 'Welcome to the PDP Sandbox. This interface is communicating via the Prompt Data Protocol. Select a privacy level and send a message to see the request headers change.' }
  ]);
  let input = $state('');
  let pdpLevel = $state(0); 
  let isSending = $state(false);

  const levels = [
    { id: 0, label: 'Private', color: 'bg-emerald-500', desc: 'No training, no storage.' },
    { id: 1, label: 'Personal', color: 'bg-amber-500', desc: 'Individual fine-tuning only.' },
    { id: 2, label: 'Global', color: 'bg-rose-500', desc: 'Full training consent.' }
  ];

  async function sendMessage(e) {
    if (e) e.preventDefault();
    if (!input.trim() || isSending) return;
    
    const userMessage = { role: 'user', content: input };
    messages = [...messages, userMessage];
    input = '';
    isSending = true;

    try {
      await fetchWithPDP('/api/chat', { method: 'POST', body: JSON.stringify({ messages }) }, pdpLevel);
      messages = [...messages, { role: 'ai', content: `PDP Request sent with Level ${pdpLevel}. The provider has been instructed to follow the ${levels[pdpLevel].label} policy.` }];
    } catch (e) {
      console.error("PDP Error:", e);
    } finally {
      isSending = false;
    }
  }
</script>

<main class="flex h-screen bg-[#F8F9FB] font-sans text-slate-900 overflow-hidden">
  <aside class="w-72 bg-white border-r border-slate-200 p-6 flex flex-col hidden lg:flex">
    <div class="flex items-center gap-3 mb-10">
      <div class="w-8 h-8 bg-black rounded-lg flex items-center justify-center">
        <span class="text-white font-bold text-xs">P</span>
      </div>
      <h1 class="font-bold text-lg tracking-tight">PDP Standard</h1>
    </div>

    <nav class="space-y-1">
      <div class="text-[10px] font-bold text-slate-400 uppercase tracking-widest mb-4">Protocol Debugger</div>
      <div class="p-4 rounded-2xl bg-slate-50 border border-slate-100 space-y-4">
        <div>
          <span class="text-[10px] font-bold text-slate-500 block mb-2 uppercase">Active Header</span>
          <code class="text-xs bg-white p-2 rounded-lg border border-slate-200 block text-blue-600">
            X-PDP-Level: {pdpLevel}
          </code>
        </div>
        <p class="text-[11px] text-slate-500 leading-relaxed italic">
          "The protocol acts as a robots.txt for AI prompts, allowing per-request consent."
        </p>
      </div>
    </nav>

    <div class="mt-auto pt-6 border-t border-slate-100 text-[10px] text-slate-400">
      v1.0.0 Draft Specification
    </div>
  </aside>

  <section class="flex-1 flex flex-col relative">
    <div class="flex-1 overflow-y-auto p-6 md:p-12 space-y-8 max-w-4xl mx-auto w-full">
      {#each messages as msg}
        <div class="flex gap-4 {msg.role === 'user' ? 'flex-row-reverse' : 'flex-row'}">
          <div class="w-8 h-8 rounded-full flex-shrink-0 flex items-center justify-center text-[10px] font-bold {msg.role === 'user' ? 'bg-black text-white' : 'bg-blue-600 text-white'}">
            {msg.role === 'user' ? 'ME' : 'AI'}
          </div>
          <div class="max-w-[85%] space-y-2 {msg.role === 'user' ? 'text-right' : 'text-left'}">
            <div class="inline-block p-4 rounded-3xl text-sm leading-relaxed shadow-sm {msg.role === 'user' ? 'bg-white border border-slate-200' : 'bg-transparent'}">
              {msg.content}
            </div>
          </div>
        </div>
      {/each}
    </div>

    <div class="p-6 max-w-4xl mx-auto w-full">
      <form onsubmit={sendMessage} class="relative bg-white border border-slate-200 rounded-[32px] shadow-xl p-2 focus-within:ring-2 focus-within:ring-slate-200 transition-all">
        
        <div class="flex gap-1 bg-slate-50 p-1 rounded-full w-fit mb-2 ml-2 mt-1">
          {#each levels as level}
            <button 
              type="button"
              onclick={() => pdpLevel = level.id}
              class="px-3 py-1 rounded-full text-[10px] font-bold transition-all flex items-center gap-1.5 {pdpLevel === level.id ? 'bg-white shadow-sm text-black' : 'text-slate-400 hover:text-slate-600'}"
            >
              <div class="w-1.5 h-1.5 rounded-full {level.color} {pdpLevel === level.id ? 'opacity-100' : 'opacity-30'}"></div>
              {level.label}
            </button>
          {/each}
        </div>

        <div class="flex items-center px-4 pb-2">
          <input 
            bind:value={input} 
            placeholder="What would you like to build today?" 
            class="flex-1 bg-transparent border-none py-3 text-sm focus:outline-none"
          />
          <button class="w-10 h-10 bg-black text-white rounded-full flex items-center justify-center hover:opacity-80 transition-all disabled:opacity-20" disabled={!input.trim()}>
            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><path d="m5 12 7-7 7 7"/><path d="M12 19V5"/></svg>
          </button>
        </div>
      </form>
      <p class="text-center text-[10px] text-slate-400 mt-4">
        PDP is an open standard. All requests in this demo are sent with explicit <strong>X-PDP-Level</strong> headers.
      </p>
    </div>
  </section>
</main>

<style>
  :global(body) {
    background: #F8F9FB;
  }
</style>
