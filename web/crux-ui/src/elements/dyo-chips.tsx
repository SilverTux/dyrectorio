import assert from 'assert'
import clsx from 'clsx'
import { useState } from 'react'

export interface DyoChipsProps<T> {
  className?: string
  key?: React.Key
  choices: readonly T[]
  initialSelection?: T
  converter?: (choice: T) => string
  onSelectionChange: (choices: T) => void
}

const DyoChips = <T,>(props: DyoChipsProps<T>) => {
  const { choices, converter, onSelectionChange, key: propsKey, className, initialSelection } = props

  assert(
    converter || choices.length < 1 || typeof choices[0] === 'string',
    'When choices are not string, you must define a converter.',
  )

  const [selection, setSelection] = useState<T>(initialSelection ?? null)

  const onToggle = (item: T) => {
    setSelection(item)
    onSelectionChange(item)
  }

  const key = propsKey ?? 'dyo-chips'
  return (
    <div className={className}>
      {choices.map((it, index) => {
        const text = converter ? converter(it) : it
        return (
          <button
            key={`${key}-${index}`}
            type="button"
            className={clsx(
              'rounded-md border-2 px-2 py-1 m-1',
              selection === it
                ? 'text-white font-medium border-dyo-turquoise bg-dyo-turquoise bg-opacity-30'
                : 'text-light-eased border-light-eased',
            )}
            onClick={() => onToggle(it)}
          >
            {text}
          </button>
        )
      })}
    </div>
  )
}

export default DyoChips
