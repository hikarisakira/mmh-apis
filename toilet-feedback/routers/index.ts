//lib/locations/index.ts
/*
編號說明

範例:
CB4A01M: 兒童醫院B4樓停車場男廁
地點-樓層-設施編號-設施類別

地點: 兒童醫院(C)、平安樓(P)、福音樓(F)
樓層: B4、B2、B1、1F，以此類推，地下室、10樓以上不帶F
設施編號: 依照Excel表格編號
設施類別: 男廁(M)、女廁(F)、無障礙廁所(D)、親子廁所(A)
*/
export const locationMap: Record<string, string> = {
    'TESTING': '資訊課測試資料',
    //兒童醫院
    //B4
    'CB4A01M': '兒童醫院B4樓停車場男廁',
    'CB4B01F': '兒童醫院B4樓停車場女廁',
    //B2
    'CB2C01A': '兒童醫院B2樓工務課旁親子廁所',
    //B1
    'CB1A01M': '兒童醫院B1樓梯廳旁男廁',
    'CB1B01F': '兒童醫院B1樓梯廳旁女廁',
    'CB1C01D': '兒童醫院B1樓梯廳旁無障礙廁所',
    //1F
    'C1FC01A': '兒童醫院1樓急診1號親子廁所',
    'C1FC02A': '兒童醫院1樓急診2號親子廁所',
    'C1FC03A': '兒童醫院1樓急診3號親子廁所',
    'C1FC04D': '兒童醫院1樓急診4號無障礙廁所',
    'C1FC05A': '兒童醫院1樓急診兒童留觀室內親子廁所',
    'C1FC06A': '兒童醫院1樓急診大廳往急診門旁親子廁所',
    'C1FC07A': '兒童醫院1樓1號電梯前親子廁所',
    'C1FA01M': '兒童醫院1樓病理檢驗科旁男廁',
    'C1FB01F': '兒童醫院1樓病理檢驗科旁女廁',
    //2F
    'C2FA01M': '兒童醫院2樓健兒門診旁男廁',
    'C2FB01F': '兒童醫院2樓健兒門診旁女廁',
    'C2FC01A': '兒童醫院2樓牙科門診旁(粉區)1號親子廁所',
    'C2FC02A': '兒童醫院2樓牙科門診旁(粉區)2號親子廁所',
    'C2FC03A': '兒童醫院2樓牙科門診旁(粉區)3號親子廁所',
    'C2FC04D': '兒童醫院2樓牙科門診旁(粉區)無障礙廁所',
    //3F
    'C3FA01M': '兒童醫院3樓兒童早療旁(藍區)男廁',
    'C3FB01F': '兒童醫院3樓兒童早療旁(藍區)女廁',
    'C3FC01A': '兒童醫院3樓兒童早療旁(藍區)1號親子廁所',
    'C3FB02F': '兒童醫院3樓中醫門診旁女廁',
    'C3FC02A': '兒童醫院3樓尿動力室旁2號親子廁所',
    'C3FC03A': '兒童醫院3樓尿動力室旁3號親子廁所',
    'C3FC04A': '兒童醫院3樓尿動力室旁4號親子廁所',
    'C3FC05D': '兒童醫院3樓尿動力室旁無障礙廁所',
    //4F
    'C4FC01A': '兒童醫院4樓家屬休息區1號親子廁所',
    'C4FC02A': '兒童醫院4樓家屬休息區2號親子廁所',
    'C4FC03D': '兒童醫院4樓產房門外無障礙廁所',
    //5F
    'C5FC01D': '兒童醫院5樓5A梯廳旁無障礙廁所',
    //6F
    'C6FC01A': '兒童醫院6樓6A病房家屬休息區旁親子廁所',
    //7F
    'C7FC01A': '兒童醫院7樓7A病房家屬休息區旁親子廁所',
    //8F
    'C8FC01A': '兒童醫院8樓8A病房家屬休息區旁親子廁所',
    //9F
    'C9FC01A': '兒童醫院9樓國際會議廳貴賓室外親子廁所',
    'C9FC02A': '兒童醫院9樓第一會議室內親子廁所',
    'C9FA01M': '兒童醫院9樓大廳男廁',
    'C9FB01F': '兒童醫院9樓大廳女廁',
    'C9FC03D': '兒童醫院9樓大廳無障礙廁所',
    //10F
    'C10A01M': '兒童醫院10樓行政辦公室內茶水間男廁',
    'C10B01F': '兒童醫院10樓行政辦公室內茶水間女廁',
    'C10A02M': '兒童醫院10樓1號電梯安全門前男廁',
    'C10B02F': '兒童醫院10樓1號電梯安全門前女廁',

    //福音樓
    //B1
    'FB1A01M': '福音樓B1樓梯廳旁男廁',
    'FB1B01F': '福音樓B1樓梯廳旁女廁',
    'FB1C01D': '福音樓B1樓梯廳旁無障礙廁所',
    'FB1C02A': '福音樓B1樓放射科旁共用廁所',
    //1F
    'F1FA01M': '福音樓1樓梯廳旁男廁',
    'F1FB01F': '福音樓1樓梯廳旁女廁',
    'F1FA02M': '福音樓1樓批價櫃台對面男廁',
    'F1FB02F': '福音樓1樓批價櫃台對面女廁',
    'F1FC01D': '福音樓1樓批價櫃台對面無障礙廁所',
    //2F
    'F2FA01M': '福音樓2樓梯廳旁男廁',
    'F2FB01F': '福音樓2樓梯廳旁女廁',
    'F2FC01A': '福音樓2樓梯廳旁共用廁所',
    'F2FB02F': '福音樓2樓癌症中心旁女廁',
    //3F
    'F3FA01A': '福音樓3樓開刀房更衣室共用廁所',
    //4F
    'F4FA01M': '福音樓4樓血液透析旁男廁',
    'F4FB01F': '福音樓4樓血液透析旁女廁',
    //6F
    'F6FA01A': '福音樓6樓病房共用廁所',
    //7F
    'F7FA01M': '福音樓7樓病房男廁',
    'F7FB01F': '福音樓7樓病房女廁',
    //8F
    'F8FA01M': '福音樓8樓病房男廁',
    'F8FB01F': '福音樓8樓病房女廁',
    //9F
    'F9FA01M': '福音樓9樓病房男廁',
    'F9FB01F': '福音樓9樓病房女廁',
    //10F
    'F10A01M': '福音樓10樓健檢中心男廁',
    'F10B01F': '福音樓10樓健檢中心女廁',

    //平安樓
    //B4
    'PB4A01M': '平安樓B4樓放腫前男廁',
    'PB4B01F': '平安樓B4樓放腫前女廁',
    //B2
    'PB2A01M': '平安樓B2樓梯廳旁男廁',
    'PB2B01F': '平安樓B2樓梯廳旁女廁',
    //1F
    'P1FA01M': '平安樓1樓梯廳旁男廁',
    'P1FB01F': '平安樓1樓梯廳旁女廁',
    //2F
    'P2FA01M': '平安樓2樓梯廳旁男廁',
    'P2FB01F': '平安樓2樓梯廳旁女廁',
    'P2FC01A': '平安樓2樓小兒科共用廁所',
    'P2FB02F': '平安樓2樓婦產科女廁',
    //3F
    'P3FA01M': '平安樓3樓梯廳旁男廁',
    'P3FB01F': '平安樓3樓梯廳旁女廁',
    //4F
    'P4FA01M': '平安樓4樓梯廳旁男廁',
    'P4FB01F': '平安樓4樓梯廳旁女廁',
    //5F
    'P5FA01M': '平安樓5樓梯廳旁男廁',
    'P5FB01F': '平安樓5樓梯廳旁女廁',
    'P5FC01A': '平安樓5樓復健科內共用廁所',
    'P5FC02A': '平安樓5樓身心科門診共用廁所',
    //6F
    'P6FA01M': '平安樓6樓36病房男廁',
    'P6FB01F': '平安樓6樓36病房女廁',
    //7F
    'P7FA01M': '平安樓7樓37病房男廁',
    'P7FB01F': '平安樓7樓37病房女廁',
    //8F
    'P8FA01M': '平安樓8樓38病房男廁',
    'P8FB01F': '平安樓8樓38病房女廁',
    //9F
    'P9FA01M': '平安樓9樓39病房男廁',
    'P9FB01F': '平安樓9樓39病房女廁',
    //10F
    'P10A01M': '平安樓10樓40病房男廁',
    'P10B01F': '平安樓10樓40病房女廁',
    //11F
    'P11A01M': '平安樓11樓41病房男廁',
    'P11B01F': '平安樓11樓41病房女廁',
    //12F
    'P12A01M': '平安樓12樓行政辦公室男廁',
    'P12B01F': '平安樓12樓行政辦公室女廁',
    //13F
    'P13A01M': '平安樓13樓國際會議廳男廁',
    'P13B01F': '平安樓13樓國際會議廳女廁',
};
