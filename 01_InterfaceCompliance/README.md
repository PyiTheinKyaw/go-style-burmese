# Verify Interface Compliance  (Interface Compliance ကို စစ်ဆေးခြင်း)

Interface Compliance ကို Compile time မှာ စစ်ဆေးသင့်ပါတယ်။ Interface Compliance ဆိုတာက Interface မှာသတ်မှတ်ထားတဲ့ Method တွေကို Derived class က Implement လုပ်ထားလား မလုပ်ထားဘူးလား ဆိုတာကို စစ်ဆေးတာဖြစ်ပါတယ်။ 

ဒီနေရာမှာ သတိထားရမှာက Go မှာ Interface တစ်ခုကို Implement လုပ်ဖို့ explicitly ကြေညာစရာမလိုပါဘူး။ ဒီအချက်ကြောင့် ကောင်းတာတွေလည်းရှိလာသလို၊ Interface ကသတ်မှတ်ထားတဲ့ Method (Behaviour) တွေများလာတဲ့အခါမျိုးဆို Implement လုပ်ဖို့မေ့တာမျိုး Method signature တွေ မကိုက်တာမျိုး ဖြစ်တတ်ပါတယ်။ (ဒီလိုဖြစ်လာရင် Go က Error ပြပေးမှာမဟုတ်ပါဘူး)

ဒါကြောင့် Interface Compliance ကိုစစ်ဆေးဖို့လိုပါတယ်။ စစ်ဆေးတဲ့နေရာမှာ အောက်ပါ အချက် ၃ ကို သေချာစိစစ်သင့်ပါတယ်။

1. Exported Types Required to Implement Specific Interfaces as Part of Their API Contract. (အပြင်ကယူသုံးမယ့် Type တွေက Interface တွေကို အကုန် Implement လုပ်ထားရမယ်)
2. Exported or unexported types that are part of a collection of types implementing the same interface. (အပြင်ကသုံးမယ့် Type တင်မကပဲ ကိုယ့် အတွင်းကပဲယူသုံးမယ့် Type တွေကလည်း Interface ကသတ်မှတ်ထားတဲ့အတိုင်း Implement လုပ်ထားရမယ်)
3. Other cases where violating an interface would break users. (အပြင်ကသုံးမယ့် Case တွေအတွက် အကုန်လုံပြီလား ထပ်စစ်ဆေးရမယ်)

### 1. Exported Types Required to Implement Specific Interfaces as Part of Their API Contract.

Exported Type ဆိုတာက package အပြင်ကနေ လှမ်းခေါ် ပြီးသုံးလို့ရတဲ့ Type အမျိုးအစားကိုပြောတာဖြစ်ပါတယ်။
API Contract ဆိုတာက Interface တစ်ခုမှာ သတ်မှတ်ထားတဲ့ Behaviour တွေကို Implement လုပ်ထားရမယ်လို့ဆိုလိုတာဖြစ်ပါတယ်။

ဆိုလိုတာက ကျွန်တော်တို့ က API သို့ library တစ်ခုကို တည်ဆောက်နေတယ်ဆိုပါစို့။

အပြင်ကို Expose လုပ်မယ့် Type တွေမှာ သတ်မှတ်ထားတဲ့ Behaviour တွေ Implement လုပ်ထားတယ်ဆိုတာကို သက်သေပြဖို့အတွက် Interface ကိုအသုံးပြုရပါတယ်။ ဆိုတော့ Interface Compliance ကို စစ်ဆေးခြင်းအားဖြင့် ကာကွယ်နိုင်ပါတယ်။ 

အကယ်လို့ Type က ကတိပေးပြီးသား Behaviour တွေကိုသာ တကယ် Implement မလုပ်ထားရင် အသုံးပြုမယ့် အပိုင်းမှာ ပြဿနာတက်မှာဖြစ်ပါတယ်။

### 2. Exported or unexported types that are part of a collection of types implementing the same interface.

Types collections တစ်ခုရဲ့ Exported or Unexported types တွေက same interface ကို Implement လုပ်ထားလား မလုပ်ထားဘူးလား ဆိုတာကိုစစ်ဆေးဖို့လိုပါတယ်။ ဥပမာ - Similar behavior but different implementation.

ဥပမာ - Shape interface ရှိမယ်။ Shape မှာက Area တွက်တဲ့ Method ရှိမယ်ပဲ ဆိုပါစို့။ Circle နဲ့ Rectangle က Shape interface ကို Implement လုပ်ထားမယ်။ ဒီလို အခြေအနေမျိုးမှာဆိုရင် Circle ကော Recentangle ကောက တကယ် Shape ကို Implement လုပ်ထားလားဆိုတဲ့ Interface Compliance ကို စစ်ဆေးသင့်ပါတယ်။

ဒီလို စစ်ဆေးခြင်းအားဖြင့် Collections of type တွေက တူညီတဲ့ Interface ကို Implement လုပ်ထားတယ်ဆိုတဲ့အချက်အပြင် implementation မှာ ဖြစ်နိုင်ချေရှိတယ် Interface implementation Inconsistency ကိုပါ ရှောင်ရှားနိုင်မှာဖြစ်ပါတယ်။

### 3. Other cases where violating an interface would break users.
Type နဲ့ Interface ကြားထဲက ကတိထားထားတဲ့ method တွေကို Implement မလုပ်မိရင် ဒါဟာ type နဲ့ interface ကြားက Contract ကို Viloate လုပ်တာပဲဖြစ်ပါတယ်။ ဒီလိုမဖြစ်ရအောင် ကျတော်တို့အနေနဲ့ Interface Compliance ကို သေချာစစ်ဆေးသင့်ပါတယ်။

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
type Handler struct {
  // ...
}



func (h *Handler) ServeHTTP(
  w http.ResponseWriter,
  r *http.Request,
) {
  ...
}
```

</td><td>

```go
type Handler struct {
  // ...
}

var _ http.Handler = (*Handler)(nil)

func (h *Handler) ServeHTTP(
  w http.ResponseWriter,
  r *http.Request,
) {
  // ...
}
```
</td></tr>
</tbody></table>

`var _ http.Handler = (*Handler)(nil)` ဆိုတဲ့ statement က Interface compliance ကိုစစ်တာဖြစ်ပါတယ်။ အကယ်လို့ Handler type က သာ `ServeHTTP` ဆိုတဲ့ Method ကို Implement မလုပ်ထားရင် Compile time မှာတင် Fail မှာဖြစ်ပြီး ကျွန်တော်တို့အနေနဲ့ ကြိုပြီး ကာကွယ်တဲ့သဘောရောက်ပါတယ်။

တကယ်က nil ကို http.Handler အမျိုးအစားအဖြစ် (*Handler) ကနေ ပြောင်းကြည့်တာဖြစ်ပါတယ်။ ပြောင်းလို့ရရင် ကျတော်တို့ Handler က http.Handler ကလိုတဲ့ 'ServeHTTP' ကို Implmement လုပ်ထားတယ်သဘောရလို့ runtime မှာမှ error တက်တာမျိုး ကို ကြိုတင်ပြီး Compile time မှာထဲက သိနိုင်တဲ့အတွက် ကြိုတင် ရှောင်ရှားနိုင်မှာဖြစ်ပါတယ်။

ဒီ Statement ရဲ့ ညာဘက်ခြမ်းက Zero value ဖြစ်ရပါမယ်။ ဥပမာ - (ခု Case ဆိုရင် *Handler အတွက် nil ကို အသုံးပြုပါတယ်), Slice တို့ Map တို့အတွက် Nil နဲ့ Struct type ဆိုရင် Empty struct ဖြစ်သင့်ပါတယ်။ (အောက်က နမူနာ ကုတ်ကိုကြည့်ပါ)

```go
type LogHandler struct {
  h   http.Handler
  log *zap.Logger
}

var _ http.Handler = LogHandler{}

func (h LogHandler) ServeHTTP(
  w http.ResponseWriter,
  r *http.Request,
) {
  // ...
}