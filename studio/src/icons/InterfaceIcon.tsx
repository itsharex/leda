import React from "react"

const path = <path d="M1015.684371 56.94439A33.916878 33.916878 0 0 0 967.781151 8.99122l-79.921951 79.921951a12.138146 12.138146 0 0 0-1.323707 1.498536 282.723902 282.723902 0 0 0-374.334439 22.228293l-79.722147 79.921951-7.992195 7.992195 1.473561 1.498537a33.891902 33.891902 0 0 0 6.518634 38.462439L784.110517 591.921951a34.01678 34.01678 0 0 0 38.51239 6.743415l1.448586 1.24878 87.88917-87.66439a282.22439 282.22439 0 0 0 22.478049-374.134634 6.643512 6.643512 0 0 0 1.523512-1.498537l80.04683-79.672195h-0.149854z m-151.801756 407.352195l-56.020293 55.945366L504.208859 216.538537l55.945365-55.945366a214.790244 214.790244 0 0 1 303.728391 303.703414zM592.972176 782.73561l-87.939122-87.66439 79.921951-79.921952a34.01678 34.01678 0 0 0 9.216-32.967805 33.916878 33.916878 0 0 0-57.194146-14.985365l-79.921952 79.921951-79.921951-79.921951 79.921951-79.921952a33.767024 33.767024 0 0 0-0.399609-47.453658 34.166634 34.166634 0 0 0-47.453659-0.499512l-79.921951 79.921951-87.939122-87.914147a33.991805 33.991805 0 0 0-38.537366-6.493658l-1.448585-1.498537-87.864195 87.914147a282.22439 282.22439 0 0 0-22.478049 374.134634c-0.524488 0.499512-1.048976 0.749268-1.523512 1.24878L9.442029 966.556098a33.916878 33.916878 0 0 0 47.978147 47.95317l79.921951-79.921951c0.549463-0.499512 0.849171-0.999024 1.323707-1.498537a282.723902 282.723902 0 0 0 374.334439-22.478048l79.921951-79.672195 0.149854-0.249757 7.992195-7.992195-1.548488-1.498536a33.492293 33.492293 0 0 0-6.443707-38.212683z m-127.925074 80.171707A214.790244 214.790244 0 1 1 161.393639 559.203902l55.945366-55.945365L520.992468 806.712195l-55.945366 55.945366v0.249756z m0 0" p-id="19329"></path>

export interface IconProps {
  size?: string | number,
  color?: string,
}

export const InerfaceSvg = <svg
  style={{ width: "16px", height: "16px" }}
  viewBox="0 0 1024 1024"
  fill="currentColor"
>
  {path}
</svg>

const InterfaceIcon = (props: IconProps) => {
  return (
    <svg
      style={{ width: props.size || "16px", height: props.size || "16px" }}
      viewBox="0 0 1024 1024"
      fill={props.color || "currentColor"}
    >
      {path}
    </svg>
  )

}
export default InterfaceIcon