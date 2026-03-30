import { TERMS_AND_CONDITIONS } from './legal'

export default function TermsConditionsPage() {
  return (
    <div className="min-h-screen bg-[#0b1220] text-gray-200">
      <div className="mx-auto max-w-4xl px-6 py-10">
        <h1 className="text-3xl font-semibold">Terms &amp; Conditions</h1>
        <div className="mt-6 whitespace-pre-wrap text-sm leading-relaxed text-white/70">
          {TERMS_AND_CONDITIONS}
        </div>
      </div>
    </div>
  )
}
